package seed

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

func (s SeedService) createPriority(items []PriorityData) error {
	var err error
	for _, item := range items {
		err = s.PriorityRepository.Create(&domain.PriorityModel{
			ID:    uuid.New(),
			Name:  item.Name,
			Color: item.Color,
		})
	}
	return err
}
