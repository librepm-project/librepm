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
	dashboards, err := c.DashboardService.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.ModelToResponseMany(*dashboards))
}

func (c DashboardController) Show(w http.ResponseWriter, r *http.Request) {
	var dashboard_id, _ = http_utils.GetParamUUID(r, "dashboard_id")
	dashboard, err := c.DashboardService.Show(dashboard_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.ModelToResponse(*dashboard))
}

func (c DashboardController) Create(w http.ResponseWriter, r *http.Request) {
	var dashboard_request DashboardRequest
	json.NewDecoder(r.Body).Decode(&dashboard_request)
	dashboard := DashboardSerializer{}.RequestToModel(dashboard_request)
	dashboard.UserID = jwt_utils.GetTokenInfoFromRequest(r).UserID
	err := c.DashboardService.Create(&dashboard)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, DashboardSerializer{}.ModelToResponse(dashboard))
}

func (c DashboardController) Update(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	var dashboard_request DashboardRequest
	json.NewDecoder(r.Body).Decode(&dashboard_request)
	dashboard := DashboardSerializer{}.RequestToModel(dashboard_request)
	err := c.DashboardService.Update(dashboard_id, &dashboard)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updated, err := c.DashboardService.Show(dashboard_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardSerializer{}.ModelToResponse(*updated))
}

func (c DashboardController) Destroy(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	err := c.DashboardService.Destroy(dashboard_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
