package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
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
	issues := c.IssueService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, IssueSerializer{}.SerializeIssues(*issues))
}

func (c IssueController) Show(w http.ResponseWriter, r *http.Request) {
	var issue_id, _ = http_utils.GetParamUUID(r, "issue_id")
	issue := c.IssueService.Show(issue_id)
	http_utils.RespondWithJSON(w, http.StatusOK, IssueSerializer{}.SerializeIssue(*issue))
}

func (c IssueController) Create(w http.ResponseWriter, r *http.Request) {
	var issue domain.IssueModel
	json.NewDecoder(r.Body).Decode(&issue)
	c.IssueService.Create(&issue)
	http_utils.RespondWithJSON(w, http.StatusCreated, IssueSerializer{}.SerializeIssue(issue))
}

func (c IssueController) Update(w http.ResponseWriter, r *http.Request) {
	issue_id, _ := http_utils.GetParamUUID(r, "issue_id")
	var issue domain.IssueModel
	json.NewDecoder(r.Body).Decode(&issue)
	c.IssueService.Update(issue_id, &issue)
	w.WriteHeader(http.StatusOK)
}

func (c IssueController) Destroy(w http.ResponseWriter, r *http.Request) {
	issue_id, _ := http_utils.GetParamUUID(r, "issue_id")
	c.IssueService.Destroy(issue_id)
	w.WriteHeader(http.StatusNoContent)
}
