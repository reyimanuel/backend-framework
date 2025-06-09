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
	app.GET("/all", t.GetAllEvent)
	app.GET("/:id", t.GetEventByID)
	app.POST("/create", middleware.MiddlewareLogin, t.CreateEvent)
	app.PATCH("/update/:id", middleware.MiddlewareLogin, t.UpdateEvent)
	app.DELETE("/delete/:id", middleware.MiddlewareLogin, t.DeleteEvent)
}

func (t *EventController) GetAllEvent(ctx *gin.Context) {
	search := ctx.Query("search")
	status := ctx.Query("status")
	sortParam := ctx.DefaultQuery("sort", "")

	response, err := t.service.GetAllEvent(search, status, sortParam)
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
	eventName := ctx.PostForm("event_name")
	eventDate := ctx.PostForm("event_date")
	eventTime := ctx.PostForm("event_time")
	eventLocation := ctx.PostForm("event_location")
	eventDescription := ctx.PostForm("event_description")
	eventOrganizer := ctx.PostForm("event_organizer")
	eventStatus := ctx.PostForm("event_status")
	if eventName == "" || eventDate == "" || eventTime == "" || eventLocation == "" || eventDescription == "" || eventOrganizer == "" || eventStatus == "" {
		HandlerError(ctx, errs.BadRequest("All fields are required"))
		return
	}

	file, err := ctx.FormFile("event_image")
	if err != nil {
		HandlerError(ctx, errs.BadRequest("Event image is required"))
		return
	}

	payload := &dto.EventRequest{
		EventName:        eventName,
		EventDate:        eventDate,
		EventTime:        eventTime,
		EventLocation:    eventLocation,
		EventDescription: eventDescription,
		EventOrganizer:   eventOrganizer,
		EventStatus:      eventStatus,
	}

	response, err := t.service.CreateEvent(ctx, payload, file)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (t *EventController) UpdateEvent(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		HandlerError(ctx, errs.BadRequest("Invalid ID"))
		return
	}

	eventName := ctx.PostForm("event_name")
	eventDate := ctx.PostForm("event_date")
	eventTime := ctx.PostForm("event_time")
	eventLocation := ctx.PostForm("event_location")
	eventDescription := ctx.PostForm("event_description")
	eventOrganizer := ctx.PostForm("event_organizer")
	eventStatus := ctx.PostForm("event_status")

	payload := &dto.EventRequest{
		EventName:        eventName,
		EventDate:        eventDate,
		EventTime:        eventTime,
		EventLocation:    eventLocation,
		EventDescription: eventDescription,
		EventOrganizer:   eventOrganizer,
		EventStatus:      eventStatus,
	}

	file, _ := ctx.FormFile("event_image")

	response, err := t.service.UpdateEvent(ctx, id, payload, file)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *EventController) DeleteEvent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		HandlerError(ctx, errs.BadRequest("Invalid ID"))
		return
	}

	response, err := t.service.DeleteEvent(uint64(id))
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
