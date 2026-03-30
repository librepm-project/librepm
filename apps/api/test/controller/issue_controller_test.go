package controller_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"
	mocksvc "apps/api/test/mock/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"libs/jwt_utils"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	os.Exit(m.Run())
}

func withIssueID(r *http.Request, id uuid.UUID) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("issue_id", id.String())
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func withKey(r *http.Request, key string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("key", key)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func withAuthToken(r *http.Request) *http.Request {
	token := jwt_utils.GenerateToken(uuid.New(), "test@example.com")
	r.Header.Set("Authorization", "Bearer "+token)
	return r
}

func buildIssueController(issueSvc domain.IssueServiceInterface) apphttp.IssueController {
	return apphttp.IssueController{
		IssueService:         issueSvc,
		IssueAuditLogService: &mocksvc.MockIssueAuditLogService{},
		NotificationService:  &mocksvc.MockNotificationService{},
		Hub:                  nil,
	}
}

func TestIssueController_Index_ReturnsAllIssues(t *testing.T) {
	issues := &[]domain.IssueModel{
		{ID: uuid.New(), Summary: "Issue A"},
		{ID: uuid.New(), Summary: "Issue B"},
	}
	svc := &mocksvc.MockIssueService{
		AllFunc: func() (*[]domain.IssueModel, error) { return issues, nil },
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/issue", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 2)
}

func TestIssueController_Index_WithFilterID_FiltersIssues(t *testing.T) {
	filterID := uuid.New()
	issues := &[]domain.IssueModel{
		{ID: uuid.New(), Summary: "Filtered Issue"},
	}
	svc := &mocksvc.MockIssueService{
		AllByFilterIDFunc: func(id uuid.UUID) (*[]domain.IssueModel, error) {
			assert.Equal(t, filterID, id)
			return issues, nil
		},
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/issue?filterId="+filterID.String(), nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Len(t, resp, 1)
}

func TestIssueController_Index_Returns400OnError(t *testing.T) {
	svc := &mocksvc.MockIssueService{
		AllFunc: func() (*[]domain.IssueModel, error) { return nil, errors.New("db error") },
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/issue", nil)
	ctrl.Index(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIssueController_Show_ReturnsIssue(t *testing.T) {
	issueID := uuid.New()
	issue := &domain.IssueModel{ID: issueID, Summary: "Test Issue"}
	svc := &mocksvc.MockIssueService{
		ShowFunc: func(id uuid.UUID) (*domain.IssueModel, error) {
			assert.Equal(t, issueID, id)
			return issue, nil
		},
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := withIssueID(
		httptest.NewRequest(http.MethodGet, "/issue/"+issueID.String(), nil),
		issueID,
	)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, issueID.String(), resp["id"])
}

func TestIssueController_Show_Returns400OnError(t *testing.T) {
	issueID := uuid.New()
	svc := &mocksvc.MockIssueService{
		ShowFunc: func(id uuid.UUID) (*domain.IssueModel, error) {
			return nil, errors.New("not found")
		},
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := withIssueID(
		httptest.NewRequest(http.MethodGet, "/issue/"+issueID.String(), nil),
		issueID,
	)
	ctrl.Show(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIssueController_ShowByKey_ReturnsIssue(t *testing.T) {
	key := "TST-1"
	issue := &domain.IssueModel{ID: uuid.New(), Key: key, Summary: "Test"}
	svc := &mocksvc.MockIssueService{
		ShowByKeyFunc: func(k string) (*domain.IssueModel, error) {
			assert.Equal(t, key, k)
			return issue, nil
		},
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := withKey(
		httptest.NewRequest(http.MethodGet, "/issue/key/"+key, nil),
		key,
	)
	ctrl.ShowByKey(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIssueController_ShowByKey_Returns404OnError(t *testing.T) {
	key := "TST-999"
	svc := &mocksvc.MockIssueService{
		ShowByKeyFunc: func(k string) (*domain.IssueModel, error) {
			return nil, errors.New("not found")
		},
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := withKey(
		httptest.NewRequest(http.MethodGet, "/issue/key/"+key, nil),
		key,
	)
	ctrl.ShowByKey(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestIssueController_Create_ReturnsCreated(t *testing.T) {
	projectID := uuid.New()
	trackerID := uuid.New()
	statusID := uuid.New()
	svc := &mocksvc.MockIssueService{
		CreateFunc: func(issue *domain.IssueModel) error { return nil },
	}
	ctrl := buildIssueController(svc)

	body := jsonBody(t, map[string]any{
		"projectId": projectID.String(),
		"trackerId": trackerID.String(),
		"statusId":  statusID.String(),
		"summary":   "New issue",
	})
	w := httptest.NewRecorder()
	r := withAuthToken(httptest.NewRequest(http.MethodPost, "/issue", body))
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestIssueController_Create_Returns400OnServiceError(t *testing.T) {
	svc := &mocksvc.MockIssueService{
		CreateFunc: func(issue *domain.IssueModel) error { return errors.New("create failed") },
	}
	ctrl := buildIssueController(svc)

	body := jsonBody(t, map[string]any{"summary": "fail"})
	w := httptest.NewRecorder()
	r := withAuthToken(httptest.NewRequest(http.MethodPost, "/issue", body))
	r.Header.Set("Content-Type", "application/json")
	ctrl.Create(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIssueController_Destroy_ReturnsNoContent(t *testing.T) {
	issueID := uuid.New()
	svc := &mocksvc.MockIssueService{
		DestroyFunc: func(id uuid.UUID) error {
			assert.Equal(t, issueID, id)
			return nil
		},
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := withIssueID(
		httptest.NewRequest(http.MethodDelete, "/issue/"+issueID.String(), nil),
		issueID,
	)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestIssueController_Destroy_Returns400OnError(t *testing.T) {
	issueID := uuid.New()
	svc := &mocksvc.MockIssueService{
		DestroyFunc: func(id uuid.UUID) error { return errors.New("delete failed") },
	}
	ctrl := buildIssueController(svc)

	w := httptest.NewRecorder()
	r := withIssueID(
		httptest.NewRequest(http.MethodDelete, "/issue/"+issueID.String(), nil),
		issueID,
	)
	ctrl.Destroy(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
