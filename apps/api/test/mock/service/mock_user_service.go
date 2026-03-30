package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockUserService struct {
	AllFunc     func() (*[]domain.UserModel, error)
	ShowFunc    func(userID uuid.UUID) (*domain.UserModel, error)
	CreateFunc  func(user *domain.UserModel) error
	UpdateFunc  func(userID uuid.UUID, user *domain.UserModel) error
	DestroyFunc func(userID uuid.UUID) error
}

func (m *MockUserService) All() (*[]domain.UserModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.UserModel{}, nil
}

func (m *MockUserService) Show(userID uuid.UUID) (*domain.UserModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(userID)
	}
	return &domain.UserModel{}, nil
}

func (m *MockUserService) Create(user *domain.UserModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(user)
	}
	return nil
}

func (m *MockUserService) Update(userID uuid.UUID, user *domain.UserModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(userID, user)
	}
	return nil
}

func (m *MockUserService) Destroy(userID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(userID)
	}
	return nil
}
