package controller

import (
	"backend/contract"
	"backend/dto"
	"backend/middleware"
	"backend/pkg/errs"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GalleryController struct {
	service contract.GalleryService
}

func (g *GalleryController) GetPrefix() string {
	return "/gallery"
}

func (g *GalleryController) InitService(service *contract.Service) {
	g.service = service.Gallery
}

func (g *GalleryController) InitRoute(app *gin.RouterGroup) {
	app.GET("/all", middleware.MiddlewareLogin, g.GetAllGalleries)
	app.GET("/:id", middleware.MiddlewareLogin, g.GetGalleryByID)
	app.POST("/create", middleware.MiddlewareLogin, g.CreateGallery)
	app.PATCH("/update/:id", middleware.MiddlewareLogin, g.UpdateGallery)
	app.DELETE("/delete/:id", middleware.MiddlewareLogin, g.DeleteGallery)
}

func (g *GalleryController) GetAllGalleries(ctx *gin.Context) {
	response, err := g.service.GetAllGalleries()
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (g *GalleryController) GetGalleryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	response, err := g.service.GetGalleryByID(idUint)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (g *GalleryController) CreateGallery(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Name is required"))
		return
	}

	description := ctx.PostForm("description")

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Image is required"))
		return
	}

	payload := dto.GalleryRequest{
		Name:        name,
		Description: description,
	}

	response, err := g.service.CreateGallery(ctx, &payload, file)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (g *GalleryController) UpdateGallery(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	var payload dto.GalleryRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, err)
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.BadRequest("Image is required"))
		return
	}

	imageName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	imagePath := fmt.Sprintf("static/%s", imageName)
	if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError("Failed to save image"))
		return
	}

	imageURL := fmt.Sprintf("/static/%s", imageName)

	response, err := g.service.UpdateGallery(idUint, &payload, imageURL)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (g *GalleryController) DeleteGallery(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	response, err := g.service.DeleteGallery(idUint)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
