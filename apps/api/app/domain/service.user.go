package domain

import "github.com/google/uuid"

type UserServiceInterface interface {
	All() (*[]UserModel, error)
	Show(user_id uuid.UUID) (*UserModel, error)
	Update(user_id uuid.UUID, user *UserModel) error
	Create(user *UserModel) error
	Destroy(user_id uuid.UUID) error
}

type UserService struct {
	UserRepository UserRepositoryInterface
}

func (s UserService) All() (*[]UserModel, error) {
	users, err := s.UserRepository.All()
	return users, err
}

func (s UserService) Create(user *UserModel) error {
	return s.UserRepository.Create(user)
}

func (s UserService) Show(user_id uuid.UUID) (*UserModel, error) {
	user, err := s.UserRepository.FindByID(user_id)
	return user, err
}

func (s UserService) Update(user_id uuid.UUID, user *UserModel) error {
	return s.UserRepository.Update(user_id, user)
}

func (s UserService) Destroy(user_id uuid.UUID) error {
	return s.UserRepository.Destroy(user_id)
}
