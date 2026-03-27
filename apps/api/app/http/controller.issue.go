package http

import (
	"net/http"

	"apps/api/app/domain"
	"github.com/google/uuid"
)

type IssueControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	ShowByKey(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type IssueController struct {
	IssueService         domain.IssueServiceInterface
	IssueAuditLogService domain.IssueAuditLogServiceInterface
}

func (c IssueController) Index(w http.ResponseWriter, r *http.Request) {
	filterIdStr := r.URL.Query().Get("filterId")
	if filterIdStr != "" {
		filterID, err := uuid.Parse(filterIdStr)
		if err != nil {
			RespondBadRequest(w)
			return
		}
		issues, err := c.IssueService.AllByFilterID(filterID)
		if err != nil {
			RespondBadRequest(w)
			return
		}
		RespondJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponseMany(*issues))
		return
	}

	issues, err := c.IssueService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponseMany(*issues))
}

func (c IssueController) Show(w http.ResponseWriter, r *http.Request) {
	issue_id, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	issue, err := c.IssueService.Show(issue_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponse(*issue))
}

func (c IssueController) ShowByKey(w http.ResponseWriter, r *http.Request) {
	key := GetParam(r, "key")
	if key == "" {
		RespondBadRequest(w)
		return
	}
	issue, err := c.IssueService.ShowByKey(key)
	if err != nil {
		RespondNotFound(w)
		return
	}
	RespondJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponse(*issue))
}

func (c IssueController) Create(w http.ResponseWriter, r *http.Request) {
	var issue_request IssueRequest
	if err := DecodeJSON(r, &issue_request); err != nil {
		RespondBadRequest(w)
		return
	}
	issue := IssueSerializer{}.RequestToModel(issue_request)
	reporterID := GetUserIDFromRequest(r)
	issue.ReporterUserID = &reporterID
	err := c.IssueService.Create(&issue)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, IssueSerializer{}.ModelToResponse(issue))
}

func (c IssueController) Update(w http.ResponseWriter, r *http.Request) {
	issue_id, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user_id := GetUserIDFromRequest(r)

	oldIssue, err := c.IssueService.Show(issue_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	var issue_request IssueRequest
	if err := DecodeJSON(r, &issue_request); err != nil {
		RespondBadRequest(w)
		return
	}
	issue := IssueSerializer{}.RequestToModel(issue_request)
	err = c.IssueService.Update(issue_id, &issue)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	updated, err := c.IssueService.Show(issue_id)
	if err != nil {
		RespondInternalError(w)
		return
	}

	c.IssueAuditLogService.LogFieldChanges(issue_id, user_id, *oldIssue, *updated)

	RespondJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponse(*updated))
}

func (c IssueController) Destroy(w http.ResponseWriter, r *http.Request) {
	issue_id, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	err = c.IssueService.Destroy(issue_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
