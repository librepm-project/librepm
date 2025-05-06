package seed

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

func (s SeedService) createStatus(items []StatusData) error {
	var err error
	for _, item := range items {
		err = s.StatusRepository.Create(&domain.StatusModel{
			ID:    uuid.New(),
			Name:  item.Name,
			Color: item.Color,
		})
	}
	return err
}
