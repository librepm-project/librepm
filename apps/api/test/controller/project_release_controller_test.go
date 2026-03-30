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

func withProjectReleaseID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("project_release_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestProjectReleaseController_Index_ReturnsOK(t *testing.T) {
	prID := uuid.New()
	releaseID := uuid.New()
	projectID := uuid.New()
	prs := &[]domain.ProjectReleaseModel{
		{
			ID:        prID,
			ReleaseID: releaseID,
			ProjectID: projectID,
		},
	}
	svc := &mocksvc.MockProjectReleaseService{
		AllFunc: func() (*[]domain.ProjectReleaseModel, error) { return prs, nil },
	}
	ctrl := apphttp.ProjectReleaseController{ProjectReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/project-release", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
}

func TestProjectReleaseController_Create_ReturnsCreated(t *testing.T) {
	releaseID := uuid.New()
	projectID := uuid.New()
	svc := &mocksvc.MockProjectReleaseService{
		CreateFunc: func(pr *domain.ProjectReleaseModel) error { return nil },
	}
	ctrl := apphttp.ProjectReleaseController{ProjectReleaseService: svc}

	body := apphttp.CreateProjectReleaseRequest{
		ReleaseID: releaseID,
		ProjectID: projectID,
	}
	bodyBytes, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/project-release", bytes.NewReader(bodyBytes))
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProjectReleaseController_Destroy_ReturnsNoContent(t *testing.T) {
	svc := &mocksvc.MockProjectReleaseService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.ProjectReleaseController{ProjectReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/project-release", nil)
	r = withProjectReleaseID(r, uuid.New())
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestProjectReleaseController_Destroy_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockProjectReleaseService{
		DestroyFunc: func(id uuid.UUID) error { return errors.New("error") },
	}
	ctrl := apphttp.ProjectReleaseController{ProjectReleaseService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/project-release", nil)
	r = withProjectReleaseID(r, uuid.New())
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
