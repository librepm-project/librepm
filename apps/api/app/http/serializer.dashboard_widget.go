package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type DashboardWidgetRequest struct {
	Type     string     `json:"type"`
	Title    string     `json:"title"`
	FilterID *uuid.UUID `json:"filterId"`
}

type DashboardWidgetResponse struct {
	ID          uuid.UUID       `json:"id"`
	DashboardID uuid.UUID       `json:"dashboardId"`
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Filter      *FilterResponse `json:"filter,omitempty"`
}

type DashboardWidgetSerializer struct{}

func (s DashboardWidgetSerializer) RequestToModel(req DashboardWidgetRequest, dashboard_id uuid.UUID) domain.DashboardWidgetModel {
	return domain.DashboardWidgetModel{
		DashboardID: dashboard_id,
		Type:        req.Type,
		Title:       req.Title,
		FilterID:    req.FilterID,
	}
}

func (s DashboardWidgetSerializer) ModelToResponse(widget domain.DashboardWidgetModel) DashboardWidgetResponse {
	resp := DashboardWidgetResponse{
		ID:          widget.ID,
		DashboardID: widget.DashboardID,
		Type:        widget.Type,
		Title:       widget.Title,
	}
	if widget.Filter != nil {
		fr := FilterSerializer{}.ModelToResponse(*widget.Filter)
		resp.Filter = &fr
	}
	return resp
}

func (s DashboardWidgetSerializer) ModelToResponseMany(widgets []domain.DashboardWidgetModel) []DashboardWidgetResponse {
	serialized := []DashboardWidgetResponse{}
	for _, w := range widgets {
		serialized = append(serialized, s.ModelToResponse(w))
	}
	return serialized
}
