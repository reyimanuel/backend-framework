package service

import (
	"backend/contract"
	"backend/dto"
	"backend/pkg/errs"
	"backend/pkg/helpers"
	"backend/pkg/token"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepository contract.AuthRepository
}

func ImplAuthService(repo *contract.Repository) contract.AuthService {
	return &AuthService{
		AuthRepository: repo.AuthRepository,
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
	user, err := a.AuthRepository.GetUserByEmail(payload.Username)
	if err != nil {
		return nil, errs.NotFound("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return nil, errs.BadRequest("invalid password")
	}

	claims := map[string]any{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	}

	accessToken, err := token.GenerateToken(claims, 3600)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err := token.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %v", err)
	}

	response := &dto.LoginResponse{
		StatusCode: 200,
		Message:    "login success",
		Data: dto.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return response, nil

}
