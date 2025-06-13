package controller

import (
	"backend/contract"
	"backend/dto"
	"backend/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service contract.AuthService
}

func (a *AuthController) GetPrefix() string {
	return "/auth"
}

func (a *AuthController) InitService(service *contract.Service) {
	a.service = service.Auth
}

func (a *AuthController) InitRoute(app *gin.RouterGroup) {
	app.POST("/login", a.login)
	app.GET("/google/login", a.GoogleLogin)
	app.GET("/google/callback", a.GoogleCallback)
}

func (a *AuthController) login(ctx *gin.Context) {
	var payload dto.LoginRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("invalid request payload"))
		return
	}

	response, err := a.service.Login(&payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (a *AuthController) GoogleLogin(ctx *gin.Context) {
	a.service.HandleGoogleLogin(ctx)
}

func (a *AuthController) GoogleCallback(ctx *gin.Context) {
	response, err := a.service.ProcessGoogleCallback(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}