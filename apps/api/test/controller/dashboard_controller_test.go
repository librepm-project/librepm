package controller_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"
	mocksvc "apps/api/test/mock/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func withDashboardID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("dashboard_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestDashboardController_Index_ReturnsOK(t *testing.T) {
	dashboards := &[]domain.DashboardModel{
		{ID: uuid.New(), Name: "Main Dashboard"},
	}
	svc := &mocksvc.MockDashboardService{
		AllFunc: func() (*[]domain.DashboardModel, error) { return dashboards, nil },
	}
	ctrl := apphttp.DashboardController{DashboardService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
	assert.Equal(t, "Main Dashboard", resp[0]["name"])
}

func TestDashboardController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockDashboardService{
		AllFunc: func() (*[]domain.DashboardModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.DashboardController{DashboardService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDashboardController_Show_ReturnsOK(t *testing.T) {
	dashboardID := uuid.New()
	svc := &mocksvc.MockDashboardService{
		ShowFunc: func(id uuid.UUID) (*domain.DashboardModel, error) {
			return &domain.DashboardModel{ID: id, Name: "KPI"}, nil
		},
	}
	ctrl := apphttp.DashboardController{DashboardService: svc}

	w := httptest.NewRecorder()
	r := withDashboardID(httptest.NewRequest(http.MethodGet, "/dashboard/"+dashboardID.String(), nil), dashboardID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, dashboardID.String(), resp["id"])
}

func TestDashboardController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockDashboardService{
		CreateFunc: func(dashboard *domain.DashboardModel) error { return nil },
	}
	ctrl := apphttp.DashboardController{DashboardService: svc}

	body := jsonBody(t, map[string]any{"name": "Sales Dashboard", "description": "Sales KPIs"})
	w := httptest.NewRecorder()
	r := withAuthToken(httptest.NewRequest(http.MethodPost, "/dashboard", body))
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "Sales Dashboard", resp["name"])
}

func TestDashboardController_Destroy_ReturnsNoContent(t *testing.T) {
	dashboardID := uuid.New()
	svc := &mocksvc.MockDashboardService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.DashboardController{DashboardService: svc}

	w := httptest.NewRecorder()
	r := withDashboardID(httptest.NewRequest(http.MethodDelete, "/dashboard/"+dashboardID.String(), nil), dashboardID)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
