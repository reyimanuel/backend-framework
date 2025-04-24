package contract

import "backend/dto"

type Service struct {
	// Add your service methods here
	Auth AuthService
}

// type exampleService interface {
// 	ExampleMethod(ctx context.Context) error
// }

type AuthService interface {
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, error)
}
