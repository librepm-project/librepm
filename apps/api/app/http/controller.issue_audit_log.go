package http

import (
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
)

type IssueAuditLogControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type IssueAuditLogController struct {
	IssueAuditLogService domain.IssueAuditLogServiceInterface
}

func (c IssueAuditLogController) Index(w http.ResponseWriter, r *http.Request) {
	issueID, err := http_utils.GetParamUUID(r, "issue_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logs, err := c.IssueAuditLogService.AllByIssue(issueID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, IssueAuditLogSerializer{}.ModelToResponseMany(*logs))
}
