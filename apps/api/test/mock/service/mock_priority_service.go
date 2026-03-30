package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockPriorityService struct {
	AllFunc     func() (*[]domain.PriorityModel, error)
	ShowFunc    func(priorityID uuid.UUID) (*domain.PriorityModel, error)
	CreateFunc  func(priority *domain.PriorityModel) error
	UpdateFunc  func(priorityID uuid.UUID, priority *domain.PriorityModel) error
	DestroyFunc func(priorityID uuid.UUID) error
}

func (m *MockPriorityService) All() (*[]domain.PriorityModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.PriorityModel{}, nil
}

func (m *MockPriorityService) Show(priorityID uuid.UUID) (*domain.PriorityModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(priorityID)
	}
	return &domain.PriorityModel{}, nil
}

func (m *MockPriorityService) Create(priority *domain.PriorityModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(priority)
	}
	return nil
}

func (m *MockPriorityService) Update(priorityID uuid.UUID, priority *domain.PriorityModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(priorityID, priority)
	}
	return nil
}

func (m *MockPriorityService) Destroy(priorityID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(priorityID)
	}
	return nil
}
