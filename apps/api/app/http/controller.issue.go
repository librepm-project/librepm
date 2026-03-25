package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"github.com/google/uuid"
)

type IssueControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type IssueController struct {
	IssueService domain.IssueServiceInterface
}

func (c IssueController) Index(w http.ResponseWriter, r *http.Request) {
	filterIdStr := r.URL.Query().Get("filterId")
	if filterIdStr != "" {
		filterID, err := uuid.Parse(filterIdStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		issues, err := c.IssueService.AllByFilterID(filterID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		http_utils.RespondWithJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponseMany(*issues))
		return
	}

	issues, err := c.IssueService.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponseMany(*issues))
}

func (c IssueController) Show(w http.ResponseWriter, r *http.Request) {
	var issue_id, _ = http_utils.GetParamUUID(r, "issue_id")
	issue, err := c.IssueService.Show(issue_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, IssueSerializer{}.ModelToResponse(*issue))
}

func (c IssueController) Create(w http.ResponseWriter, r *http.Request) {
	var issue_request IssueRequest
	json.NewDecoder(r.Body).Decode(&issue_request)
	issue := IssueSerializer{}.RequestToModel(issue_request)
	err := c.IssueService.Create(&issue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, IssueSerializer{}.ModelToResponse(issue))
}

func (c IssueController) Update(w http.ResponseWriter, r *http.Request) {
	issue_id, _ := http_utils.GetParamUUID(r, "issue_id")
	var issue_request IssueRequest
	json.NewDecoder(r.Body).Decode(&issue_request)
	issue := IssueSerializer{}.RequestToModel(issue_request)
	err := c.IssueService.Update(issue_id, &issue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c IssueController) Destroy(w http.ResponseWriter, r *http.Request) {
	issue_id, _ := http_utils.GetParamUUID(r, "issue_id")
	err := c.IssueService.Destroy(issue_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
