package http

import (
	"encoding/json"
	"net/http"
	"os"

	"apps/api/app/domain"
	"libs/http_utils"
)

type ProjectControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type ProjectController struct {
	ProjectService domain.ProjectServiceInterface
}

func (c ProjectController) Index(w http.ResponseWriter, r *http.Request) {
	projects := c.ProjectService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, ProjectSerializer{}.SerializeProjects(*projects))
}

func (c ProjectController) Show(w http.ResponseWriter, r *http.Request) {
	var project_id, _ = http_utils.GetParamUUID(r, "project_id")
	project := c.ProjectService.Show(project_id)
	http_utils.RespondWithJSON(w, http.StatusOK, ProjectSerializer{}.SerializeProject(*project))
}

func (c ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("CREATE_PRODUCT_TOKEN") != os.Getenv("CREATE_PRODUCT_TOKEN") {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var project domain.ProjectModel
	json.NewDecoder(r.Body).Decode(&project)
	c.ProjectService.Create(&project)
	http_utils.RespondWithJSON(w, http.StatusCreated, ProjectSerializer{}.SerializeProject(project))
}

func (c ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	project_id, _ := http_utils.GetParamUUID(r, "project_id")
	var project domain.ProjectModel
	json.NewDecoder(r.Body).Decode(&project)
	c.ProjectService.Update(project_id, &project)
	http_utils.RespondWithJSON(w, http.StatusOK, ProjectSerializer{}.SerializeProject(project))
}
