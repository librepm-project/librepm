package seed

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

func (s SeedService) createPriority(items []PriorityData) error {
	for _, item := range items {
		if err := s.PriorityRepository.Create(&domain.PriorityModel{
			ID:    uuid.New(),
			Name:  item.Name,
			Color: item.Color,
		}); err != nil {
			return err
		}
	}
	return nil
}
