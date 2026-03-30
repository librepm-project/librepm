package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"
	mocksvc "apps/api/test/mock/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func withReleaseID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("release_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestReleaseController_Index_ReturnsOK(t *testing.T) {
	releaseID := uuid.New()
	releases := &[]domain.ReleaseModel{
		{ID: releaseID, Name: "v1.0.0", Status: domain.ReleaseStatusPlanned},
	}
	svc := &mocksvc.MockReleaseService{
		AllFunc: func() (*[]domain.ReleaseModel, error) { return releases, nil },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/release", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
	assert.Equal(t, "v1.0.0", resp[0]["name"])
}

func TestReleaseController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockReleaseService{
		AllFunc: func() (*[]domain.ReleaseModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/release", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReleaseController_Show_ReturnsOK(t *testing.T) {
	releaseID := uuid.New()
	release := &domain.ReleaseModel{ID: releaseID, Name: "v1.0.0", Status: domain.ReleaseStatusPlanned}
	svc := &mocksvc.MockReleaseService{
		ShowFunc: func(id uuid.UUID) (*domain.ReleaseModel, error) { return release, nil },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/release", nil)
	r = withReleaseID(r, releaseID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "v1.0.0", resp["name"])
}

func TestReleaseController_Show_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockReleaseService{
		ShowFunc: func(id uuid.UUID) (*domain.ReleaseModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/release", nil)
	r = withReleaseID(r, uuid.New())
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReleaseController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockReleaseService{
		CreateFunc: func(release *domain.ReleaseModel) error { return nil },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	body := apphttp.CreateReleaseRequest{
		Name:   "v1.0.0",
		Status: domain.ReleaseStatusPlanned,
	}
	bodyBytes, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/release", bytes.NewReader(bodyBytes))
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReleaseController_Create_WithReleasedAt(t *testing.T) {
	now := time.Now()
	svc := &mocksvc.MockReleaseService{
		CreateFunc: func(release *domain.ReleaseModel) error {
			assert.NotNil(t, release.ReleasedAt)
			return nil
		},
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	body := apphttp.CreateReleaseRequest{
		Name:       "v1.0.0",
		Status:     domain.ReleaseStatusReleased,
		ReleasedAt: &now,
	}
	bodyBytes, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/release", bytes.NewReader(bodyBytes))
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestReleaseController_Update_ReturnsOK(t *testing.T) {
	releaseID := uuid.New()
	updatedRelease := &domain.ReleaseModel{
		ID:     releaseID,
		Name:   "v1.0.1",
		Status: domain.ReleaseStatusReleased,
	}
	svc := &mocksvc.MockReleaseService{
		UpdateFunc: func(id uuid.UUID, release *domain.ReleaseModel) error { return nil },
		ShowFunc:   func(id uuid.UUID) (*domain.ReleaseModel, error) { return updatedRelease, nil },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	body := apphttp.UpdateReleaseRequest{
		Name:   "v1.0.1",
		Status: domain.ReleaseStatusReleased,
	}
	bodyBytes, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/release", bytes.NewReader(bodyBytes))
	r = withReleaseID(r, releaseID)
	ctrl.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReleaseController_Destroy_ReturnsNoContent(t *testing.T) {
	svc := &mocksvc.MockReleaseService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/release", nil)
	r = withReleaseID(r, uuid.New())
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestReleaseController_Destroy_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockReleaseService{
		DestroyFunc: func(id uuid.UUID) error { return errors.New("error") },
	}
	ctrl := apphttp.ReleaseController{ReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/release", nil)
	r = withReleaseID(r, uuid.New())
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
