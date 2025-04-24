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

func (u *AuthRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
