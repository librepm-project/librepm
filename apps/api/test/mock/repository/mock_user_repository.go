package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockUserRepository struct {
	CreateFunc      func(user *domain.UserModel) error
	UpdateFunc      func(id uuid.UUID, user *domain.UserModel) error
	FindByIDFunc    func(id uuid.UUID) (*domain.UserModel, error)
	FindByEmailFunc func(email string) (*domain.UserModel, error)
	AllFunc         func() (*[]domain.UserModel, error)
	DestroyFunc     func(id uuid.UUID) error
}

func (m *MockUserRepository) Create(user *domain.UserModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(user)
	}
	return nil
}

func (m *MockUserRepository) Update(id uuid.UUID, user *domain.UserModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, user)
	}
	return nil
}

func (m *MockUserRepository) FindByID(id uuid.UUID) (*domain.UserModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.UserModel{}, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.UserModel, error) {
	if m.FindByEmailFunc != nil {
		return m.FindByEmailFunc(email)
	}
	return &domain.UserModel{}, nil
}

func (m *MockUserRepository) All() (*[]domain.UserModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.UserModel{}, nil
}

func (m *MockUserRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
