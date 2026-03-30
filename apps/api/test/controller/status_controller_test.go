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

func withStatusID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("status_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestStatusController_Index_ReturnsOK(t *testing.T) {
	statuses := &[]domain.StatusModel{
		{ID: uuid.New(), Name: "Open", Color: "#00ff00"},
	}
	svc := &mocksvc.MockStatusService{
		AllFunc: func() (*[]domain.StatusModel, error) { return statuses, nil },
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/status", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
	assert.Equal(t, "Open", resp[0]["name"])
}

func TestStatusController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockStatusService{
		AllFunc: func() (*[]domain.StatusModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/status", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestStatusController_Show_ReturnsOK(t *testing.T) {
	statusID := uuid.New()
	svc := &mocksvc.MockStatusService{
		ShowFunc: func(id uuid.UUID) (*domain.StatusModel, error) {
			return &domain.StatusModel{ID: id, Name: "Closed"}, nil
		},
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	w := httptest.NewRecorder()
	r := withStatusID(httptest.NewRequest(http.MethodGet, "/status/"+statusID.String(), nil), statusID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, statusID.String(), resp["id"])
}

func TestStatusController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockStatusService{
		CreateFunc: func(status *domain.StatusModel) error { return nil },
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	body := jsonBody(t, map[string]any{"name": "In Progress", "color": "#ffff00"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/status", body)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "In Progress", resp["name"])
}

func TestStatusController_Create_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockStatusService{
		CreateFunc: func(status *domain.StatusModel) error { return errors.New("failed") },
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	body := jsonBody(t, map[string]any{"name": "X", "color": "#000"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/status", body)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestStatusController_Update_ReturnsOK(t *testing.T) {
	statusID := uuid.New()
	svc := &mocksvc.MockStatusService{
		UpdateFunc: func(id uuid.UUID, status *domain.StatusModel) error { return nil },
		ShowFunc: func(id uuid.UUID) (*domain.StatusModel, error) {
			return &domain.StatusModel{ID: id, Name: "Updated"}, nil
		},
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	body := jsonBody(t, map[string]any{"name": "Updated", "color": "#ff0000"})
	w := httptest.NewRecorder()
	r := withStatusID(httptest.NewRequest(http.MethodPut, "/status/"+statusID.String(), body), statusID)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStatusController_Destroy_ReturnsNoContent(t *testing.T) {
	statusID := uuid.New()
	svc := &mocksvc.MockStatusService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.StatusController{StatusService: svc}

	w := httptest.NewRecorder()
	r := withStatusID(httptest.NewRequest(http.MethodDelete, "/status/"+statusID.String(), nil), statusID)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
