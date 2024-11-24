package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
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
	dashboards := c.DashboardService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.SerializeDashboards(*dashboards))
}

func (c DashboardController) Show(w http.ResponseWriter, r *http.Request) {
	var dashboard_id, _ = http_utils.GetParamUUID(r, "dashboard_id")
	dashboard := c.DashboardService.Show(dashboard_id)
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.SerializeDashboard(*dashboard))
}

func (c DashboardController) Create(w http.ResponseWriter, r *http.Request) {
	var dashboard domain.DashboardModel
	json.NewDecoder(r.Body).Decode(&dashboard)
	c.DashboardService.Create(&dashboard)
	http_utils.RespondWithJSON(w, http.StatusCreated, DashboardSerializer{}.SerializeDashboard(dashboard))
}

func (c DashboardController) Update(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	var dashboard domain.DashboardModel
	json.NewDecoder(r.Body).Decode(&dashboard)
	c.DashboardService.Update(dashboard_id, &dashboard)
	w.WriteHeader(http.StatusOK)
}

func (c DashboardController) Destroy(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	c.DashboardService.Destroy(dashboard_id)
	w.WriteHeader(http.StatusNoContent)
}
