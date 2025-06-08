package service

import (
	"backend/contract"
	"backend/dto"
	"backend/model"
	"backend/pkg/errs"
	"backend/pkg/helpers"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GalleryService struct {
	GalleryRepository contract.GalleryRepository
}

func ImplGalleryService(repo *contract.Repository) contract.GalleryService {
	return &GalleryService{
		GalleryRepository: repo.GalleryRepository,
	}
}

func (g *GalleryService) GetGalleryByID(galleryID uint64) (*dto.GalleryResponse, error) {
	gallery, err := g.GalleryRepository.GetGalleryByID(galleryID)
	if err != nil {
		return nil, errs.NotFound("Picture data not found")
	}

	data := dto.GalleryData{
		ID:          gallery.ID,
		Name:        gallery.Name,
		Description: gallery.Description,
		ImageURL:    gallery.ImageURL,
	}

	response := &dto.GalleryResponse{
		StatusCode: http.StatusOK,
		Message:    "Gallery data retrieved successfully",
		Data:       []dto.GalleryData{data},
	}

	return response, nil
}

func (g *GalleryService) GetAllGalleries() (*dto.GalleryResponse, error) {
	galleries, err := g.GalleryRepository.GetAllGalleries()
	if err != nil {
		return nil, errs.NotFound("Pictures data not found")
	}

	data := []dto.GalleryData{}
	for _, gallery := range galleries {
		data = append(data, dto.GalleryData{
			ID:          gallery.ID,
			Name:        gallery.Name,
			Description: gallery.Description,
			ImageURL:    gallery.ImageURL,
		})
	}

	response := &dto.GalleryResponse{
		StatusCode: http.StatusOK,
		Message:    "Galleries data retrieved successfully",
		Data:       data,
	}

	return response, nil
}

func (g *GalleryService) CreateGallery(ctx *gin.Context, payload *dto.GalleryRequest, file *multipart.FileHeader) (*dto.GalleryResponse, error) {
	if err := helpers.ValidateStruct(payload); err != nil {
		return nil, err
	}

	imageName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	imagePath := fmt.Sprintf("static/%s", imageName)

	if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	imageURL := fmt.Sprintf("/static/%s", imageName)

	gallery := &model.Gallery{
		Name:        payload.Name,
		Description: payload.Description,
		ImageURL:    imageURL,
	}

	newGallery, err := g.GalleryRepository.CreateGallery(gallery)
	if err != nil {
		_ = os.Remove(imagePath)
		return nil, fmt.Errorf("failed to create gallery: %w", err)
	}

	response := &dto.GalleryResponse{
		StatusCode: http.StatusCreated,
		Message:    "Gallery data created successfully",
		Data: []dto.GalleryData{
			{
				ID:          newGallery.ID,
				Name:        newGallery.Name,
				Description: newGallery.Description,
				ImageURL:    newGallery.ImageURL,
				CreatedAt:   newGallery.CreatedAt,
				UpdatedAt:   newGallery.UpdatedAt,
			},
		},
	}

	return response, nil
}

func (g *GalleryService) UpdateGallery(ctx *gin.Context, id uint64, payload *dto.GalleryRequest, file *multipart.FileHeader) (*dto.GalleryResponse, error) {
	oldGallery, err := g.GalleryRepository.GetGalleryByID(id)
	if err != nil {
		return nil, errs.NotFound("Gallery data not found")
	}

	if payload.Name == "" && payload.Description == "" && file == nil {
		return nil, errs.BadRequest("At least one field must be updated")
	}

	updatedName := oldGallery.Name
	updatedDesc := oldGallery.Description
	updatedImage := oldGallery.ImageURL

	if payload.Name != "" {
		updatedName = payload.Name
	}
	if payload.Description != "" {
		updatedDesc = payload.Description
	}

	var newImagePath, oldImagePath string
	if file != nil {
		imageName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		newImagePath = fmt.Sprintf("static/%s", imageName)
		if err := ctx.SaveUploadedFile(file, newImagePath); err != nil {
			return nil, fmt.Errorf("failed to save image: %w", err)
		}
		oldImagePath = strings.Replace(oldGallery.ImageURL, "/static/", "static/", 1)
		updatedImage = fmt.Sprintf("/static/%s", imageName)
	}

	updateData := &model.Gallery{
		Name:        updatedName,
		Description: updatedDesc,
		ImageURL:    updatedImage,
	}

	if err := g.GalleryRepository.UpdateGallery(id, updateData); err != nil {
		if file != nil {
			_ = os.Remove(newImagePath)
		}
		return nil, fmt.Errorf("failed to update gallery: %w", err)
	}

	if file != nil {
		_ = os.Remove(oldImagePath)
	}

	response := &dto.GalleryResponse{
		StatusCode: http.StatusOK,
		Message:    "Gallery data updated successfully",
		Data: []dto.GalleryData{
			{
				ID:          id,
				Name:        updateData.Name,
				Description: updateData.Description,
				ImageURL:    updateData.ImageURL,
			},
		},
	}

	return response, nil
}

func (g *GalleryService) DeleteGallery(id uint64) (*dto.GalleryResponse, error) {
	err := g.GalleryRepository.DeleteGallery(id)
	if err != nil {
		return nil, err
	}

	response := &dto.GalleryResponse{
		StatusCode: http.StatusOK,
		Message:    "Gallery data deleted successfully",
	}

	return response, nil
}
