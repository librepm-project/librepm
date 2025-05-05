package seed

import (
	"apps/api/app/domain"
	"time"
)

func (s SeedService) createFilter(items []FilterData) error {
	var err error
	for _, item := range items {
		user, _ := s.UserRepository.FindByEmail(item.UserEmail)

		filter := domain.FilterModel{
			Name:        item.Name,
			Description: item.Description,
			UserID:      user.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err = s.FilterRepository.Create(&filter)

		for _, conditionData := range item.Conditions {
			condition := domain.FilterConditionModel{
				Condition: conditionData.Condition,
				FilterID:  filter.ID,
			}

			err = s.FilterConditionRepository.Create(&condition)
		}
	}
	return err
}
