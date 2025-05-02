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
		handlerError(ctx, errs.BadRequest("invalid request payload"))
		return
	}

	response, err := a.service.Login(&payload)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
