package repository

import (
	"backend/contract"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		// Add your repository methods here
	}
}
