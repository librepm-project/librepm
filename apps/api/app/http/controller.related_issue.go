package http

import (
	"net/http"

	"apps/api/app/domain"
	"github.com/google/uuid"
)

type RelatedIssueControllerInterface interface {
	GetRelated(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type RelatedIssueController struct {
	RelatedIssueService  domain.RelatedIssueServiceInterface
	IssueAuditLogService domain.IssueAuditLogServiceInterface
}

func (c RelatedIssueController) GetRelated(w http.ResponseWriter, r *http.Request) {
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	typeFilter := r.URL.Query().Get("type")

	var typeFilterPtr *string
	if typeFilter != "" {
		typeFilterPtr = &typeFilter
	}

	relations, err := c.RelatedIssueService.GetRelatedIssues(issueID, typeFilterPtr)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	serializer := RelatedIssueSerializer{}
	responses := serializer.ModelToResponseMany(relations, issueID)

	RespondJSON(w, http.StatusOK, responses)
}

func (c RelatedIssueController) Create(w http.ResponseWriter, r *http.Request) {
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user_id := GetUserIDFromRequest(r)

	var req RelatedIssueRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}

	targetID, err := uuid.Parse(req.TargetIssueId)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	relation, err := c.RelatedIssueService.Create(issueID, targetID, req.Type)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	c.IssueAuditLogService.LogRelatedIssueAdded(issueID, user_id, *relation, relation.IssueB)

	serializer := RelatedIssueSerializer{}
	response := serializer.ModelToResponse(relation, issueID)

	RespondJSON(w, http.StatusCreated, response)
}

func (c RelatedIssueController) Delete(w http.ResponseWriter, r *http.Request) {
	relatedIssueID, err := GetParamUUID(r, "related_issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user_id := GetUserIDFromRequest(r)

	relation, err := c.RelatedIssueService.FindByID(relatedIssueID)
	if err != nil {
		RespondNotFound(w)
		return
	}

	if err := c.RelatedIssueService.Delete(relatedIssueID); err != nil {
		RespondBadRequest(w)
		return
	}

	c.IssueAuditLogService.LogRelatedIssueRemoved(relation.IssueAID, user_id, *relation, relation.IssueB)

	RespondNoContent(w)
}
