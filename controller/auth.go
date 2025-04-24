package controller

import (
	"backend/contract"
	"backend/dto"
	"backend/pkg/errs"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service contract.AuthService
}

func (a *AuthController) getPrefix() string {
	return "/auth"
}

func (a *AuthController) initService(service *contract.Service) {
	a.service = service.Auth
}

func (a *AuthController) initRoute(app *gin.RouterGroup) {
	app.POST("/login", a.login)
}

func (a *AuthController) login(ctx *gin.Context) {
	var payload dto.LoginRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errs.BadRequest("invalid request payload")
		return
	}

	result, err := a.service.Login(&payload)
	if err != nil {
		errs.BadRequest(err.Error())
		return
	}
	ctx.JSON(result.StatusCode, result)
}
