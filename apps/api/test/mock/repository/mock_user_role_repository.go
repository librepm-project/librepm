package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockUserRoleRepository struct {
	FindByUserIDFunc func(userID uuid.UUID) ([]domain.UserRoleModel, error)
	CreateFunc       func(m *domain.UserRoleModel) error
	DeleteFunc       func(userID uuid.UUID, role string) error
}

func (m *MockUserRoleRepository) FindByUserID(userID uuid.UUID) ([]domain.UserRoleModel, error) {
	if m.FindByUserIDFunc != nil {
		return m.FindByUserIDFunc(userID)
	}
	return []domain.UserRoleModel{}, nil
}

func (m *MockUserRoleRepository) Create(model *domain.UserRoleModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(model)
	}
	return nil
}

func (m *MockUserRoleRepository) Delete(userID uuid.UUID, role string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(userID, role)
	}
	return nil
}
