package domain

import (
	"libs/jwt_utils"
	"libs/password_utils"
)

type UserSessionServiceInterface interface {
	Create(email string, password string) *UserSessionCreateReturn
}

type UserSessionService struct {
	UserRepository UserRepositoryInterface
}

type UserSessionCreateReturn struct {
	User  UserModel
	Token string
}

func (s UserSessionService) Create(email string, password string) *UserSessionCreateReturn {
	user := s.UserRepository.FindByEmail(email)

	if !password_utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil
	}

	jwt_token := jwt_utils.GenerateToken(user.ID, user.Email)
	return &UserSessionCreateReturn{
		User:  *user,
		Token: jwt_token,
	}
}
