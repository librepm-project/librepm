package seed

import (
	"apps/api/app/domain"
	"time"
)

func (s SeedService) createDashboard(items []DashboardData) error {
	var err error
	for _, item := range items {
		user, _ := s.UserRepository.FindByEmail(item.UserEmail)

		err = s.DashboardRepository.Create(&domain.DashboardModel{
			Name:        item.Name,
			Description: item.Description,
			UserID:      user.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}
	return err
}
