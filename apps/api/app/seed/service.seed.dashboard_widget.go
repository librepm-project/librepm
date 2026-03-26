package seed

import (
	"apps/api/app/domain"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

func (s SeedService) createDashboardWidget(items []DashboardWidgetData) error {
	slog.Info("starting dashboard widget seeding", "count", len(items))
	for _, item := range items {
		slog.Info("processing widget", "title", item.Title, "dashboard", item.DashboardName)
		dashboard, err := s.DashboardRepository.FindByName(item.DashboardName)
		if err != nil {
			slog.Error("failed to find dashboard for widget", "dashboard_name", item.DashboardName, "error", err)
			continue
		}
		slog.Info("found dashboard", "name", dashboard.Name, "id", dashboard.ID)

		var filterID *uuid.UUID
		if item.FilterName != "" {
			filter, err := s.FilterRepository.FindByName(item.FilterName)
			if err != nil {
				slog.Error("failed to find filter for widget", "filter_name", item.FilterName, "error", err)
			} else if filter != nil {
				slog.Info("found filter", "name", filter.Name, "id", filter.ID)
				filterID = &filter.ID
			}
		}

		widget := &domain.DashboardWidgetModel{
			DashboardID: dashboard.ID,
			FilterID:    filterID,
			Type:        item.Type,
			Title:       item.Title,
			Weight:      item.Weight,
			Width:       item.Width,
			CreatedAt:   time.Now(),
		}

		err = s.DashboardWidgetRepository.Create(widget)
		if err != nil {
			slog.Error("failed to create dashboard widget", "title", item.Title, "error", err)
			return fmt.Errorf("failed to create dashboard widget '%s': %w", item.Title, err)
		}
		slog.Info("successfully created dashboard widget", "title", item.Title, "id", widget.ID)
	}
	return nil
}
