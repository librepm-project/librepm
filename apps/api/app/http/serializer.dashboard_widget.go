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

type DashboardWidgetUpdateRequest struct {
	Width *string `json:"width"`
}

type DashboardWidgetReorderItem struct {
	ID     uuid.UUID `json:"id"`
	Weight int       `json:"weight"`
}

type DashboardWidgetResponse struct {
	ID          uuid.UUID       `json:"id"`
	DashboardID uuid.UUID       `json:"dashboardId"`
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Weight      int             `json:"weight"`
	Width       string          `json:"width"`
	Filter      *FilterResponse `json:"filter,omitempty"`
}

type DashboardWidgetSerializer struct{}

func (s DashboardWidgetSerializer) RequestToModel(req DashboardWidgetRequest, dashboard_id uuid.UUID) domain.DashboardWidgetModel {
	return domain.DashboardWidgetModel{
		DashboardID: dashboard_id,
		Type:        req.Type,
		Title:       req.Title,
		FilterID:    req.FilterID,
		Width:       "1/3",
	}
}

func (s DashboardWidgetSerializer) ModelToResponse(widget domain.DashboardWidgetModel) DashboardWidgetResponse {
	width := widget.Width
	if width == "" {
		width = "1/3"
	}
	resp := DashboardWidgetResponse{
		ID:          widget.ID,
		DashboardID: widget.DashboardID,
		Type:        widget.Type,
		Title:       widget.Title,
		Weight:      widget.Weight,
		Width:       width,
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
