package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"libs/jwt_utils"
)

type DashboardControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type DashboardController struct {
	DashboardService domain.DashboardServiceInterface
}

func (c DashboardController) Index(w http.ResponseWriter, r *http.Request) {
	dashboards, _ := c.DashboardService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.ModelToResponseMany(*dashboards))
}

func (c DashboardController) Show(w http.ResponseWriter, r *http.Request) {
	var dashboard_id, _ = http_utils.GetParamUUID(r, "dashboard_id")
	dashboard, _ := c.DashboardService.Show(dashboard_id)
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.ModelToResponse(*dashboard))
}

func (c DashboardController) Create(w http.ResponseWriter, r *http.Request) {
	var dashboard_request DashboardRequest
	json.NewDecoder(r.Body).Decode(&dashboard_request)
	dashboard := DashboardSerializer{}.RequestToModel(dashboard_request)
	dashboard.UserID = jwt_utils.GetTokenInfoFromRequest(r).UserID
	c.DashboardService.Create(&dashboard)
	http_utils.RespondWithJSON(w, http.StatusCreated, DashboardSerializer{}.ModelToResponse(dashboard))
}

func (c DashboardController) Update(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	var dashboard_request DashboardRequest
	json.NewDecoder(r.Body).Decode(&dashboard_request)
	dashboard := DashboardSerializer{}.RequestToModel(dashboard_request)
	c.DashboardService.Update(dashboard_id, &dashboard)
	w.WriteHeader(http.StatusOK)
}

func (c DashboardController) Destroy(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	c.DashboardService.Destroy(dashboard_id)
	w.WriteHeader(http.StatusNoContent)
}
