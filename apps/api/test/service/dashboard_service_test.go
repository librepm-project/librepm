package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDashboardService_All_ReturnsDashboards(t *testing.T) {
	expected := &[]domain.DashboardModel{
		{ID: uuid.New(), Name: "Main Dashboard"},
	}
	repo := &mockrepo.MockDashboardRepository{
		AllFunc: func() (*[]domain.DashboardModel, error) { return expected, nil },
	}
	svc := domain.DashboardService{DashboardRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDashboardService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockDashboardRepository{
		AllFunc: func() (*[]domain.DashboardModel, error) { return nil, repoErr },
	}
	svc := domain.DashboardService{DashboardRepository: repo}

	_, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
}

func TestDashboardService_Show_ReturnsDashboard(t *testing.T) {
	dashboardID := uuid.New()
	expected := &domain.DashboardModel{ID: dashboardID, Name: "KPI"}
	repo := &mockrepo.MockDashboardRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.DashboardModel, error) {
			assert.Equal(t, dashboardID, id)
			return expected, nil
		},
	}
	svc := domain.DashboardService{DashboardRepository: repo}

	result, err := svc.Show(dashboardID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDashboardService_Create_DelegatesCorrectly(t *testing.T) {
	var saved *domain.DashboardModel
	repo := &mockrepo.MockDashboardRepository{
		CreateFunc: func(d *domain.DashboardModel) error {
			saved = d
			return nil
		},
	}
	svc := domain.DashboardService{DashboardRepository: repo}

	dashboard := &domain.DashboardModel{Name: "New Dashboard"}
	err := svc.Create(dashboard)

	assert.NoError(t, err)
	assert.Equal(t, dashboard, saved)
}

func TestDashboardService_Destroy_DelegatesCorrectly(t *testing.T) {
	dashboardID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockDashboardRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.DashboardService{DashboardRepository: repo}

	err := svc.Destroy(dashboardID)

	assert.NoError(t, err)
	assert.Equal(t, dashboardID, destroyedID)
}
