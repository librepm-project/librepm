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

func withPriorityID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("priority_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestPriorityController_Index_ReturnsOK(t *testing.T) {
	priorities := &[]domain.PriorityModel{
		{ID: uuid.New(), Name: "High", Color: "#ff0000"},
	}
	svc := &mocksvc.MockPriorityService{
		AllFunc: func() (*[]domain.PriorityModel, error) { return priorities, nil },
	}
	ctrl := apphttp.PriorityController{PriorityService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/priority", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
	assert.Equal(t, "High", resp[0]["name"])
}

func TestPriorityController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockPriorityService{
		AllFunc: func() (*[]domain.PriorityModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.PriorityController{PriorityService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/priority", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPriorityController_Show_ReturnsOK(t *testing.T) {
	priorityID := uuid.New()
	svc := &mocksvc.MockPriorityService{
		ShowFunc: func(id uuid.UUID) (*domain.PriorityModel, error) {
			return &domain.PriorityModel{ID: id, Name: "Critical"}, nil
		},
	}
	ctrl := apphttp.PriorityController{PriorityService: svc}

	w := httptest.NewRecorder()
	r := withPriorityID(httptest.NewRequest(http.MethodGet, "/priority/"+priorityID.String(), nil), priorityID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, priorityID.String(), resp["id"])
}

func TestPriorityController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockPriorityService{
		CreateFunc: func(priority *domain.PriorityModel) error { return nil },
	}
	ctrl := apphttp.PriorityController{PriorityService: svc}

	body := jsonBody(t, map[string]any{"name": "Low", "color": "#00ff00"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/priority", body)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPriorityController_Update_ReturnsOK(t *testing.T) {
	priorityID := uuid.New()
	svc := &mocksvc.MockPriorityService{
		UpdateFunc: func(id uuid.UUID, priority *domain.PriorityModel) error { return nil },
		ShowFunc: func(id uuid.UUID) (*domain.PriorityModel, error) {
			return &domain.PriorityModel{ID: id, Name: "Updated"}, nil
		},
	}
	ctrl := apphttp.PriorityController{PriorityService: svc}

	body := jsonBody(t, map[string]any{"name": "Updated", "color": "#ff0000"})
	w := httptest.NewRecorder()
	r := withPriorityID(httptest.NewRequest(http.MethodPut, "/priority/"+priorityID.String(), body), priorityID)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPriorityController_Destroy_ReturnsNoContent(t *testing.T) {
	priorityID := uuid.New()
	svc := &mocksvc.MockPriorityService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.PriorityController{PriorityService: svc}

	w := httptest.NewRecorder()
	r := withPriorityID(httptest.NewRequest(http.MethodDelete, "/priority/"+priorityID.String(), nil), priorityID)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
