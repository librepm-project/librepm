package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"github.com/google/uuid"
)

type ProjectReleaseControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	IndexByRelease(w http.ResponseWriter, r *http.Request)
	IndexByProject(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type ProjectReleaseController struct {
	ProjectReleaseService domain.ProjectReleaseServiceInterface
}

type CreateProjectReleaseRequest struct {
	ReleaseID uuid.UUID `json:"release_id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (c ProjectReleaseController) Index(w http.ResponseWriter, r *http.Request) {
	projectReleases, err := c.ProjectReleaseService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseSerializer{}.ModelToResponseMany(*projectReleases))
}

func (c ProjectReleaseController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := GetParamUUID(r, "project_release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	projectRelease, err := c.ProjectReleaseService.Show(id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseSerializer{}.ModelToResponse(*projectRelease))
}

func (c ProjectReleaseController) IndexByRelease(w http.ResponseWriter, r *http.Request) {
	releaseID, err := GetParamUUID(r, "release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	projectReleases, err := c.ProjectReleaseService.FindByReleaseID(releaseID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseSerializer{}.ModelToResponseMany(*projectReleases))
}

func (c ProjectReleaseController) IndexByProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	projectReleases, err := c.ProjectReleaseService.FindByProjectID(projectID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectReleaseSerializer{}.ModelToResponseMany(*projectReleases))
}

func (c ProjectReleaseController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectReleaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondBadRequest(w)
		return
	}

	projectRelease := domain.ProjectReleaseModel{
		ReleaseID: req.ReleaseID,
		ProjectID: req.ProjectID,
	}

	if err := c.ProjectReleaseService.Create(&projectRelease); err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusCreated, ProjectReleaseSerializer{}.ModelToResponse(projectRelease))
}

func (c ProjectReleaseController) Destroy(w http.ResponseWriter, r *http.Request) {
	id, err := GetParamUUID(r, "project_release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}

	if err := c.ProjectReleaseService.Destroy(id); err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
