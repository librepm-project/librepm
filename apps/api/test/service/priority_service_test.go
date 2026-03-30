package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPriorityService_All_ReturnsPriorities(t *testing.T) {
	expected := &[]domain.PriorityModel{
		{ID: uuid.New(), Name: "High", Color: "#ff0000"},
	}
	repo := &mockrepo.MockPriorityRepository{
		AllFunc: func() (*[]domain.PriorityModel, error) { return expected, nil },
	}
	svc := domain.PriorityService{PriorityRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPriorityService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockPriorityRepository{
		AllFunc: func() (*[]domain.PriorityModel, error) { return nil, repoErr },
	}
	svc := domain.PriorityService{PriorityRepository: repo}

	_, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
}

func TestPriorityService_Show_ReturnsPriority(t *testing.T) {
	priorityID := uuid.New()
	expected := &domain.PriorityModel{ID: priorityID, Name: "Critical"}
	repo := &mockrepo.MockPriorityRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.PriorityModel, error) {
			assert.Equal(t, priorityID, id)
			return expected, nil
		},
	}
	svc := domain.PriorityService{PriorityRepository: repo}

	result, err := svc.Show(priorityID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPriorityService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.PriorityModel
	repo := &mockrepo.MockPriorityRepository{
		CreateFunc: func(priority *domain.PriorityModel) error {
			saved = priority
			return nil
		},
	}
	svc := domain.PriorityService{PriorityRepository: repo}

	priority := &domain.PriorityModel{Name: "Low", Color: "#00ff00"}
	err := svc.Create(priority)

	assert.NoError(t, err)
	assert.Equal(t, priority, saved)
}

func TestPriorityService_Destroy_DelegatesCorrectly(t *testing.T) {
	priorityID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockPriorityRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.PriorityService{PriorityRepository: repo}

	err := svc.Destroy(priorityID)

	assert.NoError(t, err)
	assert.Equal(t, priorityID, destroyedID)
}
