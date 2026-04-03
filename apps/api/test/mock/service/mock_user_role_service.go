package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockUserRoleService struct {
	FindByUserIDFunc func(userID uuid.UUID) ([]domain.UserRoleModel, error)
	CreateFunc       func(m *domain.UserRoleModel) error
	DeleteFunc       func(userID uuid.UUID, role string) error
}

func (m *MockUserRoleService) FindByUserID(userID uuid.UUID) ([]domain.UserRoleModel, error) {
	if m.FindByUserIDFunc != nil {
		return m.FindByUserIDFunc(userID)
	}
	return []domain.UserRoleModel{}, nil
}

func (m *MockUserRoleService) Create(model *domain.UserRoleModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(model)
	}
	return nil
}

func (m *MockUserRoleService) Delete(userID uuid.UUID, role string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(userID, role)
	}
	return nil
}
