package http

import (
	"net/http"

	"apps/api/app/domain"
)

type IssueWorklogControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type IssueWorklogController struct {
	IssueWorklogService  domain.IssueWorklogServiceInterface
	IssueAuditLogService domain.IssueAuditLogServiceInterface
}

func (c IssueWorklogController) Index(w http.ResponseWriter, r *http.Request) {
	issue_id, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	worklogs, err := c.IssueWorklogService.AllByIssue(issue_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, IssueWorklogSerializer{}.ModelToResponseMany(*worklogs))
}

func (c IssueWorklogController) Create(w http.ResponseWriter, r *http.Request) {
	issue_id, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user_id := GetUserIDFromRequest(r)

	var req IssueWorklogRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}

	worklog := IssueWorklogSerializer{}.RequestToModel(req)
	worklog.IssueID = issue_id
	worklog.UserID = user_id

	err = c.IssueWorklogService.Create(&worklog)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	created, err := c.IssueWorklogService.Show(worklog.ID)
	if err != nil {
		RespondInternalError(w)
		return
	}

	c.IssueAuditLogService.LogWorklogAdded(user_id, *created)

	RespondJSON(w, http.StatusCreated, IssueWorklogSerializer{}.ModelToResponse(*created))
}

func (c IssueWorklogController) Update(w http.ResponseWriter, r *http.Request) {
	worklog_id, err := GetParamUUID(r, "worklog_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user_id := GetUserIDFromRequest(r)

	oldWorklog, err := c.IssueWorklogService.Show(worklog_id)
	if err != nil {
		RespondNotFound(w)
		return
	}

	var req IssueWorklogRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}

	worklog := IssueWorklogSerializer{}.RequestToModel(req)
	err = c.IssueWorklogService.Update(worklog_id, &worklog)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	updated, err := c.IssueWorklogService.Show(worklog_id)
	if err != nil {
		RespondInternalError(w)
		return
	}

	c.IssueAuditLogService.LogWorklogUpdated(user_id, *oldWorklog, *updated)

	RespondJSON(w, http.StatusOK, IssueWorklogSerializer{}.ModelToResponse(*updated))
}

func (c IssueWorklogController) Destroy(w http.ResponseWriter, r *http.Request) {
	worklog_id, err := GetParamUUID(r, "worklog_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user_id := GetUserIDFromRequest(r)

	worklog, err := c.IssueWorklogService.Show(worklog_id)
	if err != nil {
		RespondNotFound(w)
		return
	}

	err = c.IssueWorklogService.Destroy(worklog_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	c.IssueAuditLogService.LogWorklogRemoved(user_id, *worklog)

	RespondNoContent(w)
}
