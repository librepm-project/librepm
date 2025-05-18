package http

import (
	"apps/api/app/domain"
	"net/http"

	"libs/http_utils"
)

type ProjectIssuePropertyControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type ProjectIssuePropertyController struct {
	ProjectService domain.ProjectServiceInterface
}

func (c ProjectIssuePropertyController) Index(w http.ResponseWriter, r *http.Request) {
	var project_id, _ = http_utils.GetParamUUID(r, "project_id")
	project_issue_property, err := c.ProjectService.ShowIssuePropertyById(project_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, project_issue_property)
}
