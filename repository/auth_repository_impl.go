package repository

import (
	"echo/auth/entity"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Register(user entity.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Login(email, password string) (entity.User, error) {
	user := entity.User{}
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
