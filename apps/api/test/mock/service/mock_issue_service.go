package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockIssueService struct {
	AllFunc          func() (*[]domain.IssueModel, error)
	AllByFilterIDFunc func(filterID uuid.UUID) (*[]domain.IssueModel, error)
	ShowFunc         func(id uuid.UUID) (*domain.IssueModel, error)
	ShowByKeyFunc    func(key string) (*domain.IssueModel, error)
	CreateFunc       func(issue *domain.IssueModel) error
	UpdateFunc       func(id uuid.UUID, issue *domain.IssueModel) error
	DestroyFunc      func(id uuid.UUID) error
}

func (m *MockIssueService) All() (*[]domain.IssueModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.IssueModel{}, nil
}

func (m *MockIssueService) AllByFilterID(filterID uuid.UUID) (*[]domain.IssueModel, error) {
	if m.AllByFilterIDFunc != nil {
		return m.AllByFilterIDFunc(filterID)
	}
	return &[]domain.IssueModel{}, nil
}

func (m *MockIssueService) Show(id uuid.UUID) (*domain.IssueModel, error) {
	if m.ShowFunc != nil {
		return m.ShowFunc(id)
	}
	return &domain.IssueModel{}, nil
}

func (m *MockIssueService) ShowByKey(key string) (*domain.IssueModel, error) {
	if m.ShowByKeyFunc != nil {
		return m.ShowByKeyFunc(key)
	}
	return &domain.IssueModel{}, nil
}

func (m *MockIssueService) Create(issue *domain.IssueModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(issue)
	}
	return nil
}

func (m *MockIssueService) Update(id uuid.UUID, issue *domain.IssueModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, issue)
	}
	return nil
}

func (m *MockIssueService) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
