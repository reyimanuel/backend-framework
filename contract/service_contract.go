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
	Event   EventService
}

// type exampleService interface {
// 	ExampleMethod(ctx context.Context) error
// }

type AuthService interface {
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, error)
	ProcessGoogleCallback(ctx *gin.Context) (any, error)
	HandleGoogleLogin(c *gin.Context)
}

type TeamService interface {
	GetMemberByID(id uint64) (*dto.TeamResponse, error)
	GetAllMember(search, status, sort string) (*dto.TeamResponse, error)
	CreateMember(payload *dto.TeamRequest) (*dto.TeamResponse, error)
	UpdateMember(id uint64, payload *dto.TeamRequest) (*dto.TeamResponse, error)
	DeleteMember(id uint64) (*dto.TeamResponse, error)
}

type GalleryService interface {
	GetGalleryByID(galleryID uint64) (*dto.GalleryResponse, error)
	GetAllGalleries(search, sort string) (*dto.GalleryResponse, error)
	CreateGallery(ctx *gin.Context, payload *dto.GalleryRequest, file *multipart.FileHeader) (*dto.GalleryResponse, error)
	UpdateGallery(ctx *gin.Context, id uint64, payload *dto.GalleryRequest, file *multipart.FileHeader) (*dto.GalleryResponse, error)
	DeleteGallery(id uint64) (*dto.GalleryResponse, error)
}

type EventService interface {
	GetAllEvent(search, status, sort string) (*dto.EventResponse, error)
	GetEventByID(id uint64) (*dto.EventResponse, error)
	CreateEvent(ctx *gin.Context, payload *dto.EventRequest, file *multipart.FileHeader) (*dto.EventResponse, error)
	UpdateEvent(ctx *gin.Context, id uint64, payload *dto.EventRequest, file *multipart.FileHeader) (*dto.EventResponse, error)
	DeleteEvent(id uint64) (*dto.EventResponse, error)
}
