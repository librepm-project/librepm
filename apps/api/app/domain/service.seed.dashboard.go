package domain

import (
	"time"
)

type DashboardData struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	UserEmail   string `yaml:"user_email"`
}

func (s SeedService) createDashboard(items []DashboardData) error {
	var err error
	for _, item := range items {
		user, _ := s.UserRepository.FindByEmail(item.UserEmail)

		err = s.DashboardRepository.Create(&DashboardModel{
			Name:        item.Name,
			Description: item.Description,
			UserID:      user.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}
	return err
}
