package controller_test

import (
	"bytes"
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

func withProjectID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("project_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func jsonBody(t *testing.T, v any) *bytes.Buffer {
	t.Helper()
	b, err := json.Marshal(v)
	assert.NoError(t, err)
	return bytes.NewBuffer(b)
}

func TestProjectController_Index_ReturnsOK(t *testing.T) {
	projects := &[]domain.ProjectModel{
		{ID: uuid.New(), Name: "Alpha", CodeName: "AL"},
	}
	svc := &mocksvc.MockProjectService{
		AllFunc: func() (*[]domain.ProjectModel, error) { return projects, nil },
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/project", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
	assert.Equal(t, "Alpha", resp[0]["name"])
}

func TestProjectController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockProjectService{
		AllFunc: func() (*[]domain.ProjectModel, error) { return nil, errors.New("db error") },
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/project", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProjectController_Show_ReturnsOK(t *testing.T) {
	projectID := uuid.New()
	project := &domain.ProjectModel{ID: projectID, Name: "Alpha"}
	svc := &mocksvc.MockProjectService{
		ShowFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			assert.Equal(t, projectID, id)
			return project, nil
		},
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	w := httptest.NewRecorder()
	r := withProjectID(httptest.NewRequest(http.MethodGet, "/project/"+projectID.String(), nil), projectID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, projectID.String(), resp["id"])
}

func TestProjectController_Show_Returns400OnError(t *testing.T) {
	projectID := uuid.New()
	svc := &mocksvc.MockProjectService{
		ShowFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			return nil, errors.New("not found")
		},
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	w := httptest.NewRecorder()
	r := withProjectID(httptest.NewRequest(http.MethodGet, "/project/"+projectID.String(), nil), projectID)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProjectController_Create_ReturnsCreated(t *testing.T) {
	svc := &mocksvc.MockProjectService{
		CreateFunc: func(project *domain.ProjectModel) error { return nil },
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	body := jsonBody(t, map[string]any{
		"name":     "Beta Project",
		"codeName": "BT",
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/project", body)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "Beta Project", resp["name"])
	assert.Equal(t, "BT", resp["codeName"])
}

func TestProjectController_Create_Returns400OnServiceError(t *testing.T) {
	svc := &mocksvc.MockProjectService{
		CreateFunc: func(project *domain.ProjectModel) error { return errors.New("create failed") },
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	body := jsonBody(t, map[string]any{"name": "X", "codeName": "X"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/project", body)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProjectController_Update_ReturnsOK(t *testing.T) {
	projectID := uuid.New()
	updated := &domain.ProjectModel{ID: projectID, Name: "Updated"}
	callCount := 0
	svc := &mocksvc.MockProjectService{
		UpdateFunc: func(id uuid.UUID, project *domain.ProjectModel) error { return nil },
		ShowFunc: func(id uuid.UUID) (*domain.ProjectModel, error) {
			callCount++
			return updated, nil
		},
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	body := jsonBody(t, map[string]any{"name": "Updated", "codeName": "UP"})
	w := httptest.NewRecorder()
	r := withProjectID(
		httptest.NewRequest(http.MethodPut, "/project/"+projectID.String(), body),
		projectID,
	)
	r.Header.Set("Content-Type", "application/json")
	ctrl.Update(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, callCount)
}

func TestProjectController_Destroy_ReturnsNoContent(t *testing.T) {
	projectID := uuid.New()
	svc := &mocksvc.MockProjectService{
		DestroyFunc: func(id uuid.UUID) error {
			assert.Equal(t, projectID, id)
			return nil
		},
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	w := httptest.NewRecorder()
	r := withProjectID(
		httptest.NewRequest(http.MethodDelete, "/project/"+projectID.String(), nil),
		projectID,
	)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestProjectController_Destroy_Returns400OnError(t *testing.T) {
	projectID := uuid.New()
	svc := &mocksvc.MockProjectService{
		DestroyFunc: func(id uuid.UUID) error { return errors.New("delete failed") },
	}
	ctrl := apphttp.ProjectController{ProjectService: svc}

	w := httptest.NewRecorder()
	r := withProjectID(
		httptest.NewRequest(http.MethodDelete, "/project/"+projectID.String(), nil),
		projectID,
	)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
