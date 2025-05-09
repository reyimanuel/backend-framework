package contract

import "backend/dto"

type Service struct {
	// Add your service methods here
	Auth AuthService
	Team TeamService
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
	CreateMember(team *dto.TeamRequest) (*dto.TeamResponse, error)
	UpdateMember(id uint64, team *dto.TeamRequest) (*dto.TeamResponse, error)
	DeleteMember(id uint64) (*dto.TeamResponse, error)
}
