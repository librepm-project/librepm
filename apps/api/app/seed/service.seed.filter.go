package seed

import (
	"apps/api/app/domain"
	"time"
)

type FilterData struct {
	Name        string          `yaml:"name"`
	Description string          `yaml:"description"`
	Conditions  []ConditionData `yaml:"conditions"`
	UserEmail   string          `yaml:"user_email"`
}

type ConditionData struct {
	Condition string `yaml:"condition"`
}

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
