package contract

import "backend/model"

type Repository struct {
	AuthRepository AuthRepository
	// Add your repository methods here
}

// type exampleRepository interface {
// 	ExampleMethod(ctx context.Context) error
// }

type AuthRepository interface {
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
