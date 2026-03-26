package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"libs/jwt_utils"
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
	issue_id, _ := http_utils.GetParamUUID(r, "issue_id")
	worklogs, err := c.IssueWorklogService.AllByIssue(issue_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, IssueWorklogSerializer{}.ModelToResponseMany(*worklogs))
}

func (c IssueWorklogController) Create(w http.ResponseWriter, r *http.Request) {
	issue_id, _ := http_utils.GetParamUUID(r, "issue_id")
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID

	var req IssueWorklogRequest
	json.NewDecoder(r.Body).Decode(&req)

	worklog := IssueWorklogSerializer{}.RequestToModel(req)
	worklog.IssueID = issue_id
	worklog.UserID = user_id

	err := c.IssueWorklogService.Create(&worklog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	created, err := c.IssueWorklogService.Show(worklog.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.IssueAuditLogService.LogWorklogAdded(user_id, *created)

	http_utils.RespondWithJSON(w, http.StatusCreated, IssueWorklogSerializer{}.ModelToResponse(*created))
}

func (c IssueWorklogController) Update(w http.ResponseWriter, r *http.Request) {
	worklog_id, _ := http_utils.GetParamUUID(r, "worklog_id")
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID

	oldWorklog, err := c.IssueWorklogService.Show(worklog_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var req IssueWorklogRequest
	json.NewDecoder(r.Body).Decode(&req)

	worklog := IssueWorklogSerializer{}.RequestToModel(req)
	err = c.IssueWorklogService.Update(worklog_id, &worklog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updated, err := c.IssueWorklogService.Show(worklog_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.IssueAuditLogService.LogWorklogUpdated(user_id, *oldWorklog, *updated)

	http_utils.RespondWithJSON(w, http.StatusOK, IssueWorklogSerializer{}.ModelToResponse(*updated))
}

func (c IssueWorklogController) Destroy(w http.ResponseWriter, r *http.Request) {
	worklog_id, _ := http_utils.GetParamUUID(r, "worklog_id")
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID

	worklog, err := c.IssueWorklogService.Show(worklog_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = c.IssueWorklogService.Destroy(worklog_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c.IssueAuditLogService.LogWorklogRemoved(user_id, *worklog)

	w.WriteHeader(http.StatusNoContent)
}
