package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockTrackerService struct {
	AllFunc     func() (*[]domain.TrackerModel, error)
	ShowFunc    func(trackerID uuid.UUID) (*domain.TrackerModel, error)
	CreateFunc  func(tracker *domain.TrackerModel) error
	UpdateFunc  func(trackerID uuid.UUID, tracker *domain.TrackerModel) error
	DestroyFunc func(trackerID uuid.UUID) error
}

func (m *MockTrackerService) All() (*[]domain.TrackerModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.TrackerModel{}, nil
}

func (m *MockTrackerService) Show(trackerID uuid.UUID) (*domain.TrackerModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(trackerID)
	}
	return &domain.TrackerModel{}, nil
}

func (m *MockTrackerService) Create(tracker *domain.TrackerModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(tracker)
	}
	return nil
}

func (m *MockTrackerService) Update(trackerID uuid.UUID, tracker *domain.TrackerModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(trackerID, tracker)
	}
	return nil
}

func (m *MockTrackerService) Destroy(trackerID uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(trackerID)
	}
	return nil
}
