package domain

type TrackerData struct {
	Name string `yaml:"name"`
}

func (s SeedService) createTracker(items []TrackerData) error {
	var err error
	for _, item := range items {
		err = s.TrackerRepository.Create(&TrackerModel{
			Name: item.Name,
		})
	}
	return err
}
