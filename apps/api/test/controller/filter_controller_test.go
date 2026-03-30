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

func withFilterID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("filter_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestFilterController_Index_ReturnsOK(t *testing.T) {
	filters := &[]domain.FilterModel{
		{ID: uuid.New(), Name: "My Filter"},
	}
	svc := &mocksvc.MockFilterService{
		AllFunc: func() (*[]domain.FilterModel, error) { return filters, nil },
	}
	ctrl := apphttp.FilterController{FilterService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/filter", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
}

func TestFilterController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockFilterService{
		AllFunc: func() (*[]domain.FilterModel, error) { return nil, errors.New("error") },
	}
	ctrl := apphttp.FilterController{FilterService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/filter", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFilterController_Show_ReturnsOK(t *testing.T) {
	filterID := uuid.New()
	svc := &mocksvc.MockFilterService{
		ShowFunc: func(id uuid.UUID) (*domain.FilterModel, error) {
			return &domain.FilterModel{ID: id, Name: "Test"}, nil
		},
	}
	ctrl := apphttp.FilterController{FilterService: svc}

	w := httptest.NewRecorder()
	r := withFilterID(httptest.NewRequest(http.MethodGet, "/filter/"+filterID.String(), nil), filterID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFilterController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockFilterService{
		CreateFunc: func(filter *domain.FilterModel) error { return nil },
	}
	ctrl := apphttp.FilterController{FilterService: svc}

	body := jsonBody(t, map[string]any{"name": "New Filter", "public": false})
	w := httptest.NewRecorder()
	r := withAuthToken(httptest.NewRequest(http.MethodPost, "/filter", body))
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFilterController_Update_ReturnsOK(t *testing.T) {
	filterID := uuid.New()
	svc := &mocksvc.MockFilterService{
		UpdateFunc: func(id uuid.UUID, filter *domain.FilterModel) error { return nil },
		ShowFunc: func(id uuid.UUID) (*domain.FilterModel, error) {
			return &domain.FilterModel{ID: id, Name: "Updated"}, nil
		},
	}
	ctrl := apphttp.FilterController{FilterService: svc}

	body := jsonBody(t, map[string]any{"name": "Updated"})
	w := httptest.NewRecorder()
	r := withFilterID(httptest.NewRequest(http.MethodPut, "/filter/"+filterID.String(), body), filterID)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFilterController_Destroy_ReturnsNoContent(t *testing.T) {
	filterID := uuid.New()
	svc := &mocksvc.MockFilterService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.FilterController{FilterService: svc}

	w := httptest.NewRecorder()
	r := withFilterID(httptest.NewRequest(http.MethodDelete, "/filter/"+filterID.String(), nil), filterID)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestFilterController_Options_ReturnsFields(t *testing.T) {
	ctrl := apphttp.FilterController{FilterService: &mocksvc.MockFilterService{}}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/filter/condition-options", nil)
	ctrl.Options(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Contains(t, resp, "fields")
}
