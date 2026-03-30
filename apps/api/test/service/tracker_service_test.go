package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTrackerService_All_ReturnsTrackers(t *testing.T) {
	expected := &[]domain.TrackerModel{
		{ID: uuid.New(), Name: "Bug", Color: "#ff0000"},
	}
	repo := &mockrepo.MockTrackerRepository{
		AllFunc: func() (*[]domain.TrackerModel, error) { return expected, nil },
	}
	svc := domain.TrackerService{TrackerRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTrackerService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockTrackerRepository{
		AllFunc: func() (*[]domain.TrackerModel, error) { return nil, repoErr },
	}
	svc := domain.TrackerService{TrackerRepository: repo}

	_, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
}

func TestTrackerService_Show_ReturnsTracker(t *testing.T) {
	trackerID := uuid.New()
	expected := &domain.TrackerModel{ID: trackerID, Name: "Feature"}
	repo := &mockrepo.MockTrackerRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.TrackerModel, error) {
			assert.Equal(t, trackerID, id)
			return expected, nil
		},
	}
	svc := domain.TrackerService{TrackerRepository: repo}

	result, err := svc.Show(trackerID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestTrackerService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.TrackerModel
	repo := &mockrepo.MockTrackerRepository{
		CreateFunc: func(tracker *domain.TrackerModel) error {
			saved = tracker
			return nil
		},
	}
	svc := domain.TrackerService{TrackerRepository: repo}

	tracker := &domain.TrackerModel{Name: "Task", Color: "#0000ff"}
	err := svc.Create(tracker)

	assert.NoError(t, err)
	assert.Equal(t, tracker, saved)
}

func TestTrackerService_Update_DelegatesCorrectly(t *testing.T) {
	trackerID := uuid.New()
	var updatedID uuid.UUID
	repo := &mockrepo.MockTrackerRepository{
		UpdateFunc: func(id uuid.UUID, tracker *domain.TrackerModel) error {
			updatedID = id
			return nil
		},
	}
	svc := domain.TrackerService{TrackerRepository: repo}

	err := svc.Update(trackerID, &domain.TrackerModel{Name: "Updated"})

	assert.NoError(t, err)
	assert.Equal(t, trackerID, updatedID)
}

func TestTrackerService_Destroy_DelegatesCorrectly(t *testing.T) {
	trackerID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockTrackerRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.TrackerService{TrackerRepository: repo}

	err := svc.Destroy(trackerID)

	assert.NoError(t, err)
	assert.Equal(t, trackerID, destroyedID)
}
