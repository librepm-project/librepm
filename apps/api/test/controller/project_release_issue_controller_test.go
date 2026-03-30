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

func withProjectReleaseIssueID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("project_release_issue_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestProjectReleaseIssueController_Index_ReturnsOK(t *testing.T) {
	priID := uuid.New()
	prIssues := &[]domain.ProjectReleaseIssueModel{
		{
			ID:               priID,
			ProjectReleaseID: uuid.New(),
			IssueID:          uuid.New(),
		},
	}
	svc := &mocksvc.MockProjectReleaseIssueService{
		AllFunc: func() (*[]domain.ProjectReleaseIssueModel, error) { return prIssues, nil },
	}
	ctrl := apphttp.ProjectReleaseIssueController{ProjectReleaseIssueService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/project-release-issue", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
}

func TestProjectReleaseIssueController_Create_ReturnsCreated(t *testing.T) {
	prID := uuid.New()
	issueID := uuid.New()
	svc := &mocksvc.MockProjectReleaseIssueService{
		CreateFunc: func(pri *domain.ProjectReleaseIssueModel) error { return nil },
	}
	ctrl := apphttp.ProjectReleaseIssueController{ProjectReleaseIssueService: svc}

	body := apphttp.CreateProjectReleaseIssueRequest{
		ProjectReleaseID: prID,
		IssueID:          issueID,
	}
	bodyBytes, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/project-release-issue", bytes.NewReader(bodyBytes))
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProjectReleaseIssueController_Destroy_ReturnsNoContent(t *testing.T) {
	svc := &mocksvc.MockProjectReleaseIssueService{
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := apphttp.ProjectReleaseIssueController{ProjectReleaseIssueService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/project-release-issue", nil)
	r = withProjectReleaseIssueID(r, uuid.New())
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestProjectReleaseIssueController_Destroy_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockProjectReleaseIssueService{
		DestroyFunc: func(id uuid.UUID) error { return errors.New("error") },
	}
	ctrl := apphttp.ProjectReleaseIssueController{ProjectReleaseIssueService: svc}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/project-release-issue", nil)
	r = withProjectReleaseIssueID(r, uuid.New())
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
