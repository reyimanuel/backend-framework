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

func (a *AuthRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func (a *AuthRepository) FindOrCreateUser(email, username string) (any, error) {
	var user model.User
	result := a.db.Where("email = ?", email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		user = model.User{
			Email:    email,
			Username: username,
			Password: "",
		}
		if err := a.db.Create(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
	}
	return user, result.Error
}