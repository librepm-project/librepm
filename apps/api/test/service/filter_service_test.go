package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFilterService_All_ReturnsFilters(t *testing.T) {
	expected := &[]domain.FilterModel{
		{ID: uuid.New(), Name: "Filter A"},
		{ID: uuid.New(), Name: "Filter B"},
	}
	filterRepo := &mockrepo.MockFilterRepository{
		AllFunc: func() (*[]domain.FilterModel, error) {
			return expected, nil
		},
	}
	svc := domain.FilterService{FilterRepository: filterRepo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestFilterService_All_ReturnsErrorOnFailure(t *testing.T) {
	repoErr := errors.New("db error")
	filterRepo := &mockrepo.MockFilterRepository{
		AllFunc: func() (*[]domain.FilterModel, error) {
			return nil, repoErr
		},
	}
	svc := domain.FilterService{FilterRepository: filterRepo}

	result, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
	assert.Nil(t, result)
}

func TestFilterService_Show_ReturnsFilter(t *testing.T) {
	filterID := uuid.New()
	expected := &domain.FilterModel{ID: filterID, Name: "My Filter"}
	filterRepo := &mockrepo.MockFilterRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.FilterModel, error) {
			assert.Equal(t, filterID, id)
			return expected, nil
		},
	}
	svc := domain.FilterService{FilterRepository: filterRepo}

	result, err := svc.Show(filterID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestFilterService_Create_DelegatesCorrectly(t *testing.T) {
	var savedFilter *domain.FilterModel
	filterRepo := &mockrepo.MockFilterRepository{
		CreateFunc: func(filter *domain.FilterModel) error {
			savedFilter = filter
			return nil
		},
	}
	svc := domain.FilterService{FilterRepository: filterRepo}

	filter := &domain.FilterModel{Name: "New Filter", Public: true}
	err := svc.Create(filter)

	assert.NoError(t, err)
	assert.Equal(t, filter, savedFilter)
}

func TestFilterService_Update_DelegatesCorrectly(t *testing.T) {
	filterID := uuid.New()
	var updatedID uuid.UUID
	var updatedFilter *domain.FilterModel

	filterRepo := &mockrepo.MockFilterRepository{
		UpdateFunc: func(id uuid.UUID, filter *domain.FilterModel) error {
			updatedID = id
			updatedFilter = filter
			return nil
		},
	}
	svc := domain.FilterService{FilterRepository: filterRepo}

	filter := &domain.FilterModel{Name: "Updated Filter"}
	err := svc.Update(filterID, filter)

	assert.NoError(t, err)
	assert.Equal(t, filterID, updatedID)
	assert.Equal(t, filter, updatedFilter)
}

func TestFilterService_Destroy_DelegatesCorrectly(t *testing.T) {
	filterID := uuid.New()
	var destroyedID uuid.UUID

	filterRepo := &mockrepo.MockFilterRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.FilterService{FilterRepository: filterRepo}

	err := svc.Destroy(filterID)

	assert.NoError(t, err)
	assert.Equal(t, filterID, destroyedID)
}
