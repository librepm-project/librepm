package domain

import (
	"libs/password_utils"
)

type UserRegisterServiceInterface interface {
	Create(email string, password string, firstName string, lastName string) (*UserModel, error)
}

type UserRegisterService struct {
	UserRepository     UserRepositoryInterface
	UserRoleRepository UserRoleRepositoryInterface
}

type UserRegisterCreateReturn struct {
	User  UserModel
	Token string
}

func (s UserRegisterService) Create(email string, password string, firstName string, lastName string) (*UserModel, error) {
	password_hash, _ := password_utils.HashPassword(password)
	user := UserModel{
		Email:        email,
		PasswordHash: password_hash,
		FirstName:    firstName,
		LastName:     lastName,
	}
	if err := s.UserRepository.Create(&user); err != nil {
		return &user, err
	}
	defaultRole := "viewer"
	_ = s.UserRoleRepository.Create(&UserRoleModel{
		UserID: user.ID,
		Role:   defaultRole,
	})
	return &user, nil
}
