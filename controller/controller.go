package controller

import (
	"backend/contract"
	"backend/middleware"
	"backend/pkg/errs"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller interface {
	// getPrefix returns the route prefix that the controller will use.
	getPrefix() string

	// initService initializes the necessary services for the controller.
	// This service typically contains the business logic required by the controller.
	initService(service *contract.Service)

	// initRoute sets up the routes for the controller within the given router group.
	initRoute(app *gin.RouterGroup)
}

func New(app *gin.Engine, service *contract.Service) {
	allController := []controller{
		// Add your controller here
		&AuthController{},
	}

	// do not modify the code below there
	for _, c := range allController {
		c.initService(service)
		group := app.Group(c.getPrefix())
		group.Use(middleware.CORSMiddleware())
		c.initRoute(group)
		log.Printf("initiate route %s\n", c.getPrefix())
	}
}

// handlerError is a helper function to handle errors in the controller.
// It checks if the error is of type MessageError and responds with the appropriate status code and message.
func handlerError(ctx *gin.Context, err error) {
	var messageErr errs.MessageError
	if errors.As(err, &messageErr) {
		ctx.JSON(messageErr.Status(), messageErr)
	} else {
		ctx.Error(err).SetType(gin.ErrorTypePrivate)
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError("Internal Server Error"))
	}
}
