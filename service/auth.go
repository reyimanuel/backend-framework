package service

import (
	"backend/config"
	"backend/contract"
	"backend/dto"
	"backend/model"
	"backend/pkg/errs"
	"backend/pkg/helpers"
	"backend/pkg/token"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	AuthRepository contract.AuthRepository
	OAuthConfig    *oauth2.Config
}

func ImplAuthService(repo *contract.Repository) contract.AuthService {
	config := config.Get()
	return &AuthService{
		AuthRepository: repo.AuthRepository,
		OAuthConfig: &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.RedirectURL,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint: google.Endpoint,
		},
	}
}

func (a *AuthService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, validPayload
	}

	var missingFields []string
	if payload.Username == "" {
		missingFields = append(missingFields, "username")
	}

	if payload.Password == "" {
		missingFields = append(missingFields, "password")
	}
	if len(missingFields) > 0 {
		return nil, errs.BadRequest("missing required fields: " + strings.Join(missingFields, ", "))
	}
	user, err := a.AuthRepository.GetUserByUsername(payload.Username)
	if err != nil {
		return nil, errs.Unauthorized("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return nil, errs.Unauthorized("invalid username or password")
	}

	accessToken, err := token.GenerateToken(&token.UserAuthToken{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	log.Printf("Access Token: %s", accessToken)

	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	response := &dto.LoginResponse{
		StatusCode: http.StatusOK,
		Message:    "login success",
		Data: dto.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return response, nil

}

func (a *AuthService) HandleGoogleLogin(c *gin.Context) {
	url := a.OAuthConfig.AuthCodeURL("random-state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (a *AuthService) ProcessGoogleCallback(ctx *gin.Context) (any, error) {
	code := ctx.Query("code")
	if code == "" {
		return nil, errs.BadRequest("missing code in URL")
	}

	// Exchange code untuk dapat token
	tokenResp, err := a.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	client := a.OAuthConfig.Client(ctx, tokenResp)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info response: %v", err)
	}

	email, ok1 := userInfo["email"].(string)
	name, ok2 := userInfo["name"].(string)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("invalid user info response: %v", userInfo)
	}

	// Find or Create User di DB
	user, err := a.AuthRepository.FindOrCreateUser(email, name)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create user: %v", err)
	}

	users, ok := user.(model.User)
	if !ok {
		return nil, fmt.Errorf("unexpected user type: %T", user)
	}

	accessToken, err := token.GenerateToken(&token.UserAuthToken{
		ID:       users.ID,
		Email:    users.Email,
		Username: users.Username,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}
	refreshToken, err := token.GenerateRefreshToken(users.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}
	response := &dto.GoogleAuthResponse{
		StatusCode: http.StatusOK,
		Message:    "Google login successful",
		Data: map[string]interface{}{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}

	return response, nil
}