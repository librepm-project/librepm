package domain

import "github.com/google/uuid"

type UserServiceInterface interface {
	All() *[]UserModel
	Show(user_id uuid.UUID) *UserModel
	Update(user_id uuid.UUID, user *UserModel)
	Create(user *UserModel)
	Destroy(user_id uuid.UUID)
}

type UserService struct {
	UserRepository UserRepositoryInterface
}

func (s UserService) All() *[]UserModel {
	users := s.UserRepository.All()
	return users
}

func (s UserService) Create(user *UserModel) {
	s.UserRepository.Create(user)
}

func (s UserService) Show(user_id uuid.UUID) *UserModel {
	user := s.UserRepository.FindByID(user_id)
	return user
}

func (s UserService) Update(user_id uuid.UUID, user *UserModel) {
	s.UserRepository.Update(user_id, user)
}

func (s UserService) Destroy(user_id uuid.UUID) {
	s.UserRepository.Destroy(user_id)
}
