package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockIssueRepository struct {
	AllFunc         func() (*[]domain.IssueModel, error)
	AllByFilterFunc func(conditions []domain.FilterConditionModel) (*[]domain.IssueModel, error)
	FindByIDFunc    func(id uuid.UUID) (*domain.IssueModel, error)
	FindByKeyFunc   func(key string) (*domain.IssueModel, error)
	CreateFunc      func(issue *domain.IssueModel) error
	UpdateFunc      func(id uuid.UUID, issue *domain.IssueModel) error
	DestroyFunc     func(id uuid.UUID) error
}

func (m *MockIssueRepository) All() (*[]domain.IssueModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.IssueModel{}, nil
}

func (m *MockIssueRepository) AllByFilter(conditions []domain.FilterConditionModel) (*[]domain.IssueModel, error) {
	if m.AllByFilterFunc != nil {
		return m.AllByFilterFunc(conditions)
	}
	return &[]domain.IssueModel{}, nil
}

func (m *MockIssueRepository) FindByID(id uuid.UUID) (*domain.IssueModel, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return &domain.IssueModel{}, nil
}

func (m *MockIssueRepository) FindByKey(key string) (*domain.IssueModel, error) {
	if m.FindByKeyFunc != nil {
		return m.FindByKeyFunc(key)
	}
	return &domain.IssueModel{}, nil
}

func (m *MockIssueRepository) Create(issue *domain.IssueModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(issue)
	}
	return nil
}

func (m *MockIssueRepository) Update(id uuid.UUID, issue *domain.IssueModel) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(id, issue)
	}
	return nil
}

func (m *MockIssueRepository) Destroy(id uuid.UUID) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(id)
	}
	return nil
}
