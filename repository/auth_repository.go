package repository

import "echo/auth/entity"

type AuthRepository interface {
	Register(user entity.User) error
	Login(email, password string) (entity.User, error)
}
