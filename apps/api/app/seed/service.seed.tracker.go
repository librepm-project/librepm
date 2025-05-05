package seed

import "apps/api/app/domain"

func (s SeedService) createTracker(items []TrackerData) error {
	var err error
	for _, item := range items {
		err = s.TrackerRepository.Create(&domain.TrackerModel{
			Name: item.Name,
		})
	}
	return err
}
