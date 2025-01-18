package domain

import (
	"libs/password_utils"
)

type UserRegisterServiceInterface interface {
	Create(email string, password string) (*UserModel, error)
}

type UserRegisterService struct {
	UserRepository UserRepositoryInterface
}

type UserRegisterCreateReturn struct {
	User  UserModel
	Token string
}

func (s UserRegisterService) Create(email string, password string) (*UserModel, error) {
	password_hash, _ := password_utils.HashPassword(password)
	user := UserModel{
		Email:        email,
		PasswordHash: password_hash,
	}
	err := s.UserRepository.Create(&user)
	return &user, err
}
