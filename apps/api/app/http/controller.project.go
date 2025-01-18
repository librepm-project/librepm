package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
)

type ProjectControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type ProjectController struct {
	ProjectService domain.ProjectServiceInterface
}

func (c ProjectController) Index(w http.ResponseWriter, r *http.Request) {
	projects, _ := c.ProjectService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, ProjectSerializer{}.ModelToResponseMany(*projects))
}

func (c ProjectController) Show(w http.ResponseWriter, r *http.Request) {
	var project_id, _ = http_utils.GetParamUUID(r, "project_id")
	project, _ := c.ProjectService.Show(project_id)
	http_utils.RespondWithJSON(w, http.StatusOK, ProjectSerializer{}.ModelToResponse(*project))
}

func (c ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	var project_request ProjectRequest
	json.NewDecoder(r.Body).Decode(&project_request)
	project := ProjectSerializer{}.RequestToModel(project_request)
	c.ProjectService.Create(&project)
	http_utils.RespondWithJSON(w, http.StatusCreated, ProjectSerializer{}.ModelToResponse(project))
}

func (c ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	project_id, _ := http_utils.GetParamUUID(r, "project_id")
	var project_request ProjectRequest
	json.NewDecoder(r.Body).Decode(&project_request)
	project := ProjectSerializer{}.RequestToModel(project_request)
	c.ProjectService.Update(project_id, &project)
	w.WriteHeader(http.StatusOK)
}

func (c ProjectController) Destroy(w http.ResponseWriter, r *http.Request) {
	project_id, _ := http_utils.GetParamUUID(r, "project_id")
	c.ProjectService.Destroy(project_id)
	w.WriteHeader(http.StatusNoContent)
}
