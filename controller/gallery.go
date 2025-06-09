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
	search := ctx.Query("search")
	sort := ctx.DefaultQuery("sort", "") // contoh: created_at:desc

	response, err := g.service.GetAllGalleries(search, sort)
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

	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	file, _ := ctx.FormFile("image")

	payload := dto.GalleryRequest{
		Name:        name,
		Description: description,
	}

	response, err := g.service.UpdateGallery(ctx, idUint, &payload, file)
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
