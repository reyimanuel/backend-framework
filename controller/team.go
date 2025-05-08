package controller

import (
	"backend/contract"
	"backend/dto"
	"backend/middleware"
	"backend/pkg/errs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	service contract.TeamService
}

func (t *TeamController) GetPrefix() string {
	return "/team"
}

func (t *TeamController) InitService(service *contract.Service) {
	t.service = service.Team
}

func (t *TeamController) InitRoute(app *gin.RouterGroup) {
	app.GET("/all", middleware.MiddlewareLogin, t.GetAllMember)
	app.GET("/:id", middleware.MiddlewareLogin, t.GetMemberByID)
	app.POST("/create", middleware.MiddlewareLogin, t.CreateMember)
	app.PATCH("/update/:id", middleware.MiddlewareLogin, t.UpdateMember)
	app.DELETE("/delete/:id", middleware.MiddlewareLogin, t.DeleteMember)
	app.GET("/division/:division", middleware.MiddlewareLogin, t.GetMemberByDivision)
}

func (t *TeamController) GetAllMember(ctx *gin.Context) {
	response, err := t.service.GetAllMember()
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *TeamController) GetMemberByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	response, err := t.service.GetMemberByID(idUint)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *TeamController) GetMemberByDivision(ctx *gin.Context) {
	division := ctx.Param("division")
	response, err := t.service.GetMemberByDivision(division)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *TeamController) CreateMember(ctx *gin.Context) {
	var payload dto.TeamRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("invalid request payload"))
		return
	}
	response, err := t.service.CreateMember(&payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *TeamController) UpdateMember(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	var payload dto.TeamRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("invalid request payload"))
		return
	}
	response, err := t.service.UpdateMember(idUint, &payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *TeamController) DeleteMember(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	response, err := t.service.DeleteMember(idUint)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
