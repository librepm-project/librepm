package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"github.com/google/uuid"
)

type ProjectReleaseIssueControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	IndexByProjectRelease(w http.ResponseWriter, r *http.Request)
	ShowByIssue(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type ProjectReleaseIssueController struct {
	ProjectReleaseIssueService domain.ProjectReleaseIssueServiceInterface
}

type CreateProjectReleaseIssueRequest struct {
	ProjectReleaseID uuid.UUID `json:"project_release_id"`
	IssueID          uuid.UUID `json:"issue_id"`
}

func (c ProjectReleaseIssueController) Index(w http.ResponseWriter, r *http.Request) {
	prIssues, err := c.ProjectReleaseIssueService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseIssueSerializer{}.ModelToResponseMany(*prIssues))
}

func (c ProjectReleaseIssueController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := GetParamUUID(r, "project_release_issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	prIssue, err := c.ProjectReleaseIssueService.Show(id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseIssueSerializer{}.ModelToResponse(*prIssue))
}

func (c ProjectReleaseIssueController) IndexByProjectRelease(w http.ResponseWriter, r *http.Request) {
	projectReleaseID, err := GetParamUUID(r, "project_release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	prIssues, err := c.ProjectReleaseIssueService.FindByProjectReleaseID(projectReleaseID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseIssueSerializer{}.ModelToResponseMany(*prIssues))
}

func (c ProjectReleaseIssueController) ShowByIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	prIssue, err := c.ProjectReleaseIssueService.FindByIssueID(issueID)
	if err != nil {
		RespondJSON(w, http.StatusOK, nil)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseIssueSerializer{}.ModelToResponse(*prIssue))
}

func (c ProjectReleaseIssueController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectReleaseIssueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondBadRequest(w)
		return
	}

	prIssue := domain.ProjectReleaseIssueModel{
		ProjectReleaseID: req.ProjectReleaseID,
		IssueID:          req.IssueID,
	}

	if err := c.ProjectReleaseIssueService.Create(&prIssue); err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusCreated, ProjectReleaseIssueSerializer{}.ModelToResponse(prIssue))
}

func (c ProjectReleaseIssueController) Destroy(w http.ResponseWriter, r *http.Request) {
	id, err := GetParamUUID(r, "project_release_issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}

	if err := c.ProjectReleaseIssueService.Destroy(id); err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
