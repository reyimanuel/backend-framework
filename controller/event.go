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

type EventController struct {
	service contract.EventService
}

func (t *EventController) GetPrefix() string {
	return "/event"
}

func (t *EventController) InitService(service *contract.Service) {
	t.service = service.Event
}

func (t *EventController) InitRoute(app *gin.RouterGroup) {
	app.GET("/all", middleware.MiddlewareLogin, t.GetAllEvent)
	app.GET("/:id", middleware.MiddlewareLogin, t.GetEventByID)
	app.POST("/create", middleware.MiddlewareLogin, t.CreateEvent)
}

func (t *EventController) GetAllEvent(ctx *gin.Context) {
	response, err := t.service.GetAllEvent()
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *EventController) GetEventByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	response, err := t.service.GetEventByID(idUint)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (t *EventController) CreateEvent(ctx *gin.Context) {
	var payload dto.EventRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("invalid request payload"))
		return
	}
	response, err := t.service.CreateEvent(&payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
