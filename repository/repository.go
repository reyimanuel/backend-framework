package repository

import (
	"backend/contract"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
		AuthRepository:    ImplAuthRepository(db),
		TeamRepository:    ImplTeamRepository(db),
		GalleryRepository: ImplGalleryRepository(db),
		// Add your repository methods here
	}
}
