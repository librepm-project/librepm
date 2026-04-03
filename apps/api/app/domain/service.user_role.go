package domain

import "github.com/google/uuid"

type UserRoleServiceInterface interface {
	FindByUserID(userID uuid.UUID) ([]UserRoleModel, error)
	Create(m *UserRoleModel) error
	Delete(userID uuid.UUID, role string) error
}

type UserRoleService struct {
	UserRoleRepository UserRoleRepositoryInterface
}

func (s UserRoleService) FindByUserID(userID uuid.UUID) ([]UserRoleModel, error) {
	return s.UserRoleRepository.FindByUserID(userID)
}

func (s UserRoleService) Create(m *UserRoleModel) error {
	return s.UserRoleRepository.Create(m)
}

func (s UserRoleService) Delete(userID uuid.UUID, role string) error {
	return s.UserRoleRepository.Delete(userID, role)
}
