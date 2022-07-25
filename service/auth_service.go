package service

import "echo/auth/entity"

type AuthService interface {
	Register(user entity.User) error
	Login(email, password string) (entity.User, error)
}
