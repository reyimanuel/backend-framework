package repository

import (
	"backend/contract"
	"backend/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func ImplAuthRepository(db *gorm.DB) contract.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (u *AuthRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
