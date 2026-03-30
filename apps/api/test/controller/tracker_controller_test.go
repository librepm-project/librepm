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

func withTrackerID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("tracker_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestTrackerController_Index_ReturnsOK(t *testing.T) {
	trackers := &[]domain.TrackerModel{
		{ID: uuid.New(), Name: "Bug", Color: "#ff0000"},
	}
	svc := &mocksvc.MockTrackerService{
		AllFunc: func() (*[]domain.TrackerModel, error) { return trackers, nil },
	}
	ctrl := apphttp.TrackerController{TrackerService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/tracker", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
	assert.Equal(t, "Bug", resp[0]["name"])
}

func TestTrackerController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockTrackerService{
		AllFunc: func() (*[]domain.TrackerModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.TrackerController{TrackerService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/tracker", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTrackerController_Show_ReturnsOK(t *testing.T) {
	trackerID := uuid.New()
	svc := &mocksvc.MockTrackerService{
		ShowFunc: func(id uuid.UUID) (*domain.TrackerModel, error) {
			return &domain.TrackerModel{ID: id, Name: "Feature"}, nil
		},
	}
	ctrl := apphttp.TrackerController{TrackerService: svc}

	w := httptest.NewRecorder()
	r := withTrackerID(httptest.NewRequest(http.MethodGet, "/tracker/"+trackerID.String(), nil), trackerID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, trackerID.String(), resp["id"])
}

func TestTrackerController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockTrackerService{
		CreateFunc: func(tracker *domain.TrackerModel) error { return nil },
	}
	ctrl := apphttp.TrackerController{TrackerService: svc}

	body := jsonBody(t, map[string]any{"name": "Task", "color": "#0000ff"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/tracker", body)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestTrackerController_Update_ReturnsOK(t *testing.T) {
	trackerID := uuid.New()
	svc := &mocksvc.MockTrackerService{
		UpdateFunc: func(id uuid.UUID, tracker *domain.TrackerModel) error { return nil },
		ShowFunc: func(id uuid.UUID) (*domain.TrackerModel, error) {
			return &domain.TrackerModel{ID: id, Name: "Updated"}, nil
		},
	}
	ctrl := apphttp.TrackerController{TrackerService: svc}

	body := jsonBody(t, map[string]any{"name": "Updated", "color": "#ff0000"})
	w := httptest.NewRecorder()
	r := withTrackerID(httptest.NewRequest(http.MethodPut, "/tracker/"+trackerID.String(), body), trackerID)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTrackerController_Destroy_ReturnsNoContent(t *testing.T) {
	trackerID := uuid.New()
	svc := &mocksvc.MockTrackerService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.TrackerController{TrackerService: svc}

	w := httptest.NewRecorder()
	r := withTrackerID(httptest.NewRequest(http.MethodDelete, "/tracker/"+trackerID.String(), nil), trackerID)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
