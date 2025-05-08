package contract

import "backend/model"

type Repository struct {
	AuthRepository AuthRepository
	TeamRepository TeamRepository
	// Add your repository methods here
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
