package mocksvc

import (
	"apps/api/app/domain"
)

type MockUserSessionService struct {
	CreateFunc func(email string, password string) (*domain.UserSessionCreateReturn, error)
}

func (m *MockUserSessionService) Create(email string, password string) (*domain.UserSessionCreateReturn, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(email, password)
	}
	return &domain.UserSessionCreateReturn{}, nil
}
