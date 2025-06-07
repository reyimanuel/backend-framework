package contract

import (
	"backend/dto"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type Service struct {
	// Add your service methods here
	Auth    AuthService
	Team    TeamService
	Gallery GalleryService
}

// type exampleService interface {
// 	ExampleMethod(ctx context.Context) error
// }

type AuthService interface {
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, error)
}

type TeamService interface {
	GetMemberByID(id uint64) (*dto.TeamResponse, error)
	GetAllMember() (*dto.TeamResponse, error)
	GetMemberByDivision(division string) (*dto.TeamResponse, error)
	CreateMember(payload *dto.TeamRequest) (*dto.TeamResponse, error)
	UpdateMember(id uint64, payload *dto.TeamRequest) (*dto.TeamResponse, error)
	DeleteMember(id uint64) (*dto.TeamResponse, error)
}

type GalleryService interface {
	GetGalleryByID(galleryID uint64) (*dto.GalleryResponse, error)
	GetAllGalleries() (*dto.GalleryResponse, error)
	CreateGallery(ctx *gin.Context, payload *dto.GalleryRequest, file *multipart.FileHeader) (*dto.GalleryResponse, error)
	UpdateGallery(id uint64, payload *dto.GalleryRequest, imageURL string) (*dto.GalleryResponse, error)
	DeleteGallery(id uint64) (*dto.GalleryResponse, error)
}
