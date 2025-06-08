package contract

import "backend/model"

type Repository struct {
	AuthRepository    AuthRepository
	TeamRepository    TeamRepository
	EventRepository   EventRepository
	GalleryRepository GalleryRepository
}

// type exampleRepository interface {
// 	ExampleMethod(ctx context.Context) error
// }

type AuthRepository interface {
	GetUserByUsername(username string) (*model.User, error)
}

type TeamRepository interface {
	GetMemberByID(id uint64) (*model.Team, error)
	GetAllMember() ([]model.Team, error)
	CreateMember(team *model.Team) (*model.Team, error)
	UpdateMember(id uint64, team *model.Team) error
	DeleteMember(id uint64) error
	GetMemberByDivision(division string) ([]model.Team, error)
}

type GalleryRepository interface {
	GetAllGalleries() ([]model.Gallery, error)
	GetGalleryByID(id uint64) (*model.Gallery, error)
	CreateGallery(gallery *model.Gallery) (*model.Gallery, error)
	UpdateGallery(id uint64, gallery *model.Gallery) error
	DeleteGallery(id uint64) error
}
type EventRepository interface {
	GetAllEvent() ([]model.Event, error)
	GetEventByID(id uint64) (*model.Event, error)
	CreateEvent(event *model.Event) (*model.Event, error)
}
