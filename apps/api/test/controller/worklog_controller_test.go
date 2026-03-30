package controller_test

import (
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

func withWorklogID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("worklog_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func withIssueIDAndWorklogID(r *http.Request, issueID, worklogID uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("issue_id", issueID.String())
	rctx.URLParams.Add("worklog_id", worklogID.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func buildWorklogController(worklogSvc domain.IssueWorklogServiceInterface) apphttp.IssueWorklogController {
	return apphttp.IssueWorklogController{
		IssueWorklogService:  worklogSvc,
		IssueAuditLogService: &mocksvc.MockIssueAuditLogService{},
	}
}

func TestWorklogController_Index_ReturnsOK(t *testing.T) {
	issueID := uuid.New()
	worklogs := &[]domain.IssueWorklogModel{
		{ID: uuid.New(), Seconds: 3600, Comment: "Work done"},
	}
	svc := &mocksvc.MockIssueWorklogService{
		AllByIssueFunc: func(id uuid.UUID) (*[]domain.IssueWorklogModel, error) {
			return worklogs, nil
		},
	}
	ctrl := buildWorklogController(svc)

	w := httptest.NewRecorder()
	r := withIssueID(httptest.NewRequest(http.MethodGet, "/issue/"+issueID.String()+"/worklog", nil), issueID)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
}

func TestWorklogController_Index_Returns400OnError(t *testing.T) {
	issueID := uuid.New()
	svc := &mocksvc.MockIssueWorklogService{
		AllByIssueFunc: func(id uuid.UUID) (*[]domain.IssueWorklogModel, error) {
			return nil, errors.New("db error")
		},
	}
	ctrl := buildWorklogController(svc)

	w := httptest.NewRecorder()
	r := withIssueID(httptest.NewRequest(http.MethodGet, "/issue/"+issueID.String()+"/worklog", nil), issueID)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWorklogController_Create_ReturnsCreated(t *testing.T) {
	issueID := uuid.New()
	worklogID := uuid.New()
	svc := &mocksvc.MockIssueWorklogService{
		CreateFunc: func(worklog *domain.IssueWorklogModel) error {
			worklog.ID = worklogID
			return nil
		},
		ShowFunc: func(id uuid.UUID) (*domain.IssueWorklogModel, error) {
			return &domain.IssueWorklogModel{ID: id, IssueID: issueID, Seconds: 3600}, nil
		},
	}
	ctrl := buildWorklogController(svc)

	body := jsonBody(t, map[string]any{
		"seconds":  3600,
		"comment":  "Worked on it",
		"loggedAt": time.Now().Format(time.RFC3339),
	})
	w := httptest.NewRecorder()
	r := withAuthToken(withIssueID(httptest.NewRequest(http.MethodPost, "/issue/"+issueID.String()+"/worklog", body), issueID))
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, float64(3600), resp["seconds"])
}

func TestWorklogController_Destroy_ReturnsNoContent(t *testing.T) {
	issueID := uuid.New()
	worklogID := uuid.New()
	svc := &mocksvc.MockIssueWorklogService{
		ShowFunc: func(id uuid.UUID) (*domain.IssueWorklogModel, error) {
			return &domain.IssueWorklogModel{ID: id, IssueID: issueID}, nil
		},
		DestroyFunc: func(id uuid.UUID) error { return nil },
	}
	ctrl := buildWorklogController(svc)

	w := httptest.NewRecorder()
	r := withAuthToken(withIssueIDAndWorklogID(
		httptest.NewRequest(http.MethodDelete, "/issue/"+issueID.String()+"/worklog/"+worklogID.String(), nil),
		issueID, worklogID,
	))
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
