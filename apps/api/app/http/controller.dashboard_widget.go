package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
)

type DashboardWidgetControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Reorder(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type DashboardWidgetController struct {
	DashboardWidgetService domain.DashboardWidgetServiceInterface
}

func (c DashboardWidgetController) Index(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	widgets, err := c.DashboardWidgetService.AllByDashboard(dashboard_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, DashboardWidgetSerializer{}.ModelToResponseMany(*widgets))
}

func (c DashboardWidgetController) Create(w http.ResponseWriter, r *http.Request) {
	dashboard_id, _ := http_utils.GetParamUUID(r, "dashboard_id")
	var req DashboardWidgetRequest
	json.NewDecoder(r.Body).Decode(&req)
	widget := DashboardWidgetSerializer{}.RequestToModel(req, dashboard_id)
	err := c.DashboardWidgetService.Create(&widget)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, DashboardWidgetSerializer{}.ModelToResponse(widget))
}

func (c DashboardWidgetController) Update(w http.ResponseWriter, r *http.Request) {
	widget_id, _ := http_utils.GetParamUUID(r, "widget_id")
	var req DashboardWidgetUpdateRequest
	json.NewDecoder(r.Body).Decode(&req)
	fields := map[string]interface{}{}
	if req.Width != nil {
		fields["width"] = *req.Width
	}
	if len(fields) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.DashboardWidgetService.Update(widget_id, fields); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c DashboardWidgetController) Reorder(w http.ResponseWriter, r *http.Request) {
	var items []DashboardWidgetReorderItem
	json.NewDecoder(r.Body).Decode(&items)
	updates := make([]domain.WeightUpdate, len(items))
	for i, item := range items {
		updates[i] = domain.WeightUpdate{ID: item.ID, Weight: item.Weight}
	}
	if err := c.DashboardWidgetService.BatchUpdateWeights(updates); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c DashboardWidgetController) Destroy(w http.ResponseWriter, r *http.Request) {
	widget_id, _ := http_utils.GetParamUUID(r, "widget_id")
	err := c.DashboardWidgetService.Destroy(widget_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
