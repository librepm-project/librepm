package mocksvc

import (
	"apps/api/app/domain"
)

type MockUserRegisterService struct {
	CreateFunc func(email string, password string, firstName string, lastName string) (*domain.UserModel, error)
}

func (m *MockUserRegisterService) Create(email string, password string, firstName string, lastName string) (*domain.UserModel, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(email, password, firstName, lastName)
	}
	return &domain.UserModel{}, nil
}
