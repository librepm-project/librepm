package http

import (
	"log/slog"
	"net/http"

	"apps/api/app/domain"
)

type ProjectControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	IndexTrackers(w http.ResponseWriter, r *http.Request)
	IndexStatuses(w http.ResponseWriter, r *http.Request)
}

type ProjectController struct {
	ProjectService domain.ProjectServiceInterface
}

func (c ProjectController) Index(w http.ResponseWriter, r *http.Request) {
	projects, err := c.ProjectService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectSerializer{}.ModelToResponseMany(*projects))
}

func (c ProjectController) Show(w http.ResponseWriter, r *http.Request) {
	project_id, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	project, err := c.ProjectService.Show(project_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectSerializer{}.ModelToResponse(*project))
}

func (c ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	var project_request ProjectRequest
	if err := DecodeJSON(r, &project_request); err != nil {
		RespondBadRequest(w)
		return
	}
	project := ProjectSerializer{}.RequestToModel(project_request)
	err := c.ProjectService.Create(&project)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, ProjectSerializer{}.ModelToResponse(project))
}

func (c ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	project_id, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var project_request ProjectRequest
	if err := DecodeJSON(r, &project_request); err != nil {
		RespondBadRequest(w)
		return
	}
	project := ProjectSerializer{}.RequestToModel(project_request)
	err = c.ProjectService.Update(project_id, &project)
	if err != nil {
		slog.Error("failed to update project", "error", err)
		RespondBadRequest(w)
		return
	}
	updated, err := c.ProjectService.Show(project_id)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, ProjectSerializer{}.ModelToResponse(*updated))
}

func (c ProjectController) IndexTrackers(w http.ResponseWriter, r *http.Request) {
	project_id, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	trackers, err := c.ProjectService.TrackersByProject(project_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, TrackerSerializer{}.ModelToResponseMany(*trackers))
}

func (c ProjectController) IndexStatuses(w http.ResponseWriter, r *http.Request) {
	project_id, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	statuses, err := c.ProjectService.StatusesByProject(project_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponseMany(*statuses))
}

func (c ProjectController) Destroy(w http.ResponseWriter, r *http.Request) {
	project_id, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	err = c.ProjectService.Destroy(project_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
