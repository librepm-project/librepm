package seed

import "apps/api/app/domain"

func NewSeedService(domain domain.Domain) SeedService {

	return SeedService{
		ProjectRepository:              domain.ProjectRepository,
		ProjectTrackerRepository:       domain.ProjectTrackerRepository,
		TrackerRepository:              domain.TrackerRepository,
		StatusRepository:               domain.StatusRepository,
		BoardRepository:                domain.BoardRepository,
		UserRepository:                 domain.UserRepository,
		DashboardRepository:            domain.DashboardRepository,
		FilterRepository:               domain.FilterRepository,
		FilterConditionRepository:      domain.FilterConditionRepository,
		ProjectTrackerStatusRepository: domain.ProjectTrackerStatusRepository,
		PurgeRepository: PurgeRepository{
			DB: domain.DB,
		},
	}
}
