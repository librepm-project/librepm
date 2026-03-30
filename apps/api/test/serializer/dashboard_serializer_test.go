package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDashboardSerializer_RequestToModel_MapsFields(t *testing.T) {
	req := apphttp.DashboardRequest{
		Name:        "Sales KPIs",
		Description: "Sales metrics",
		Public:      true,
	}

	model := apphttp.DashboardSerializer{}.RequestToModel(req)

	assert.Equal(t, "Sales KPIs", model.Name)
	assert.Equal(t, "Sales metrics", model.Description)
	assert.True(t, model.Public)
}

func TestDashboardSerializer_ModelToResponse_MapsFields(t *testing.T) {
	id := uuid.New()
	model := domain.DashboardModel{
		ID:          id,
		Name:        "Main Dashboard",
		Description: "Primary view",
		Public:      false,
	}

	resp := apphttp.DashboardSerializer{}.ModelToResponse(model)

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "Main Dashboard", resp.Name)
	assert.Equal(t, "Primary view", resp.Description)
	assert.False(t, resp.Public)
	assert.Empty(t, resp.Widgets)
}

func TestDashboardSerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.DashboardModel{
		{ID: uuid.New(), Name: "Dashboard A"},
		{ID: uuid.New(), Name: "Dashboard B"},
	}

	resp := apphttp.DashboardSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, "Dashboard A", resp[0].Name)
}

func TestDashboardSerializer_ModelToResponseMany_EmptySlice(t *testing.T) {
	resp := apphttp.DashboardSerializer{}.ModelToResponseMany([]domain.DashboardModel{})

	assert.Empty(t, resp)
}
