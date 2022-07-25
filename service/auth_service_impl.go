package service

import (
	"echo/auth/config"
	"echo/auth/entity"
	"echo/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r    repository.AuthRepository
	conf config.DBConfig
}

func NewAuthService(r repository.AuthRepository, conf config.DBConfig) AuthService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) Register(user entity.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hash)
	err := s.r.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Login(email, password string) (entity.User, error) {
	user, _ := s.r.Login(email, password)
	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return user, er
	}
	return user, nil
}
