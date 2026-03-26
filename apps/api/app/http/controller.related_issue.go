package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"github.com/google/uuid"
)

type RelatedIssueControllerInterface interface {
	GetRelated(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type RelatedIssueController struct {
	RelatedIssueService domain.RelatedIssueServiceInterface
}

func (c RelatedIssueController) GetRelated(w http.ResponseWriter, r *http.Request) {
	issueID, _ := http_utils.GetParamUUID(r, "issue_id")
	typeFilter := r.URL.Query().Get("type")

	var typeFilterPtr *string
	if typeFilter != "" {
		typeFilterPtr = &typeFilter
	}

	relations, err := c.RelatedIssueService.GetRelatedIssues(issueID, typeFilterPtr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serializer := RelatedIssueSerializer{}
	responses := serializer.ModelToResponseMany(relations, issueID)

	http_utils.RespondWithJSON(w, http.StatusOK, responses)
}

func (c RelatedIssueController) Create(w http.ResponseWriter, r *http.Request) {
	issueID, _ := http_utils.GetParamUUID(r, "issue_id")

	var req RelatedIssueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	targetID, err := uuid.Parse(req.TargetIssueId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	relation, err := c.RelatedIssueService.Create(issueID, targetID, req.Type)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serializer := RelatedIssueSerializer{}
	response := serializer.ModelToResponse(relation, issueID)

	http_utils.RespondWithJSON(w, http.StatusCreated, response)
}

func (c RelatedIssueController) Delete(w http.ResponseWriter, r *http.Request) {
	relatedIssueID, _ := http_utils.GetParamUUID(r, "related_issue_id")

	if err := c.RelatedIssueService.Delete(relatedIssueID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
