package service_test

import (
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDashboardWidgetService_AllByDashboard_ReturnsWidgets(t *testing.T) {
	dashboardID := uuid.New()
	expected := &[]domain.DashboardWidgetModel{
		{ID: uuid.New(), Type: "filter"},
	}
	repo := &mockrepo.MockDashboardWidgetRepository{
		AllByDashboardFunc: func(id uuid.UUID) (*[]domain.DashboardWidgetModel, error) {
			assert.Equal(t, dashboardID, id)
			return expected, nil
		},
	}
	svc := domain.DashboardWidgetService{DashboardWidgetRepository: repo}

	result, err := svc.AllByDashboard(dashboardID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDashboardWidgetService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.DashboardWidgetModel
	repo := &mockrepo.MockDashboardWidgetRepository{
		CreateFunc: func(widget *domain.DashboardWidgetModel) error {
			saved = widget
			return nil
		},
	}
	svc := domain.DashboardWidgetService{DashboardWidgetRepository: repo}

	widget := &domain.DashboardWidgetModel{Type: "filter"}
	err := svc.Create(widget)

	assert.NoError(t, err)
	assert.Equal(t, widget, saved)
}

func TestDashboardWidgetService_Update_DelegatesFields(t *testing.T) {
	widgetID := uuid.New()
	var updatedID uuid.UUID
	var updatedFields map[string]interface{}

	repo := &mockrepo.MockDashboardWidgetRepository{
		UpdateFunc: func(id uuid.UUID, fields map[string]interface{}) error {
			updatedID = id
			updatedFields = fields
			return nil
		},
	}
	svc := domain.DashboardWidgetService{DashboardWidgetRepository: repo}

	fields := map[string]interface{}{"title": "My Widget"}
	err := svc.Update(widgetID, fields)

	assert.NoError(t, err)
	assert.Equal(t, widgetID, updatedID)
	assert.Equal(t, fields, updatedFields)
}

func TestDashboardWidgetService_BatchUpdateWeights_DelegatesCorrectly(t *testing.T) {
	var receivedUpdates []domain.WeightUpdate
	repo := &mockrepo.MockDashboardWidgetRepository{
		BatchUpdateWeightsFunc: func(updates []domain.WeightUpdate) error {
			receivedUpdates = updates
			return nil
		},
	}
	svc := domain.DashboardWidgetService{DashboardWidgetRepository: repo}

	updates := []domain.WeightUpdate{
		{ID: uuid.New(), Weight: 1},
		{ID: uuid.New(), Weight: 2},
	}
	err := svc.BatchUpdateWeights(updates)

	assert.NoError(t, err)
	assert.Equal(t, updates, receivedUpdates)
}

func TestDashboardWidgetService_Destroy_DelegatesCorrectly(t *testing.T) {
	widgetID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockDashboardWidgetRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.DashboardWidgetService{DashboardWidgetRepository: repo}

	err := svc.Destroy(widgetID)

	assert.NoError(t, err)
	assert.Equal(t, widgetID, destroyedID)
}
