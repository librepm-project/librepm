package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransitionService_All_ReturnsTransitions(t *testing.T) {
	expected := &[]domain.TransitionModel{
		{ID: uuid.New()},
	}
	repo := &mockrepo.MockTransitionRepository{
		AllFunc: func() (*[]domain.TransitionModel, error) { return expected, nil },
	}
	svc := domain.TransitionService{TransitionRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTransitionService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockTransitionRepository{
		AllFunc: func() (*[]domain.TransitionModel, error) { return nil, repoErr },
	}
	svc := domain.TransitionService{TransitionRepository: repo}

	_, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
}

func TestTransitionService_Show_ReturnsTransition(t *testing.T) {
	transitionID := uuid.New()
	expected := &domain.TransitionModel{ID: transitionID}
	repo := &mockrepo.MockTransitionRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.TransitionModel, error) {
			assert.Equal(t, transitionID, id)
			return expected, nil
		},
	}
	svc := domain.TransitionService{TransitionRepository: repo}

	result, err := svc.Show(transitionID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTransitionService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.TransitionModel
	repo := &mockrepo.MockTransitionRepository{
		CreateFunc: func(t *domain.TransitionModel) error {
			saved = t
			return nil
		},
	}
	svc := domain.TransitionService{TransitionRepository: repo}

	transition := &domain.TransitionModel{}
	err := svc.Create(transition)

	assert.NoError(t, err)
	assert.Equal(t, transition, saved)
}

func TestTransitionService_Destroy_DelegatesCorrectly(t *testing.T) {
	transitionID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockTransitionRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.TransitionService{TransitionRepository: repo}

	err := svc.Destroy(transitionID)

	assert.NoError(t, err)
	assert.Equal(t, transitionID, destroyedID)
}
