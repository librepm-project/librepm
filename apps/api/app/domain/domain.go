package domain

import (
	"libs/mysql_utils"
)

type Domain struct {
	ProjectService      ProjectServiceInterface
	IssueService        IssueServiceInterface
	FilterService       FilterServiceInterface
	BoardService        BoardServiceInterface
	DashboardService    DashboardServiceInterface
	UserService         UserService
	UserSessionService  UserSessionService
	UserRegisterService UserRegisterService
	StatusService       StatusServiceInterface
	TransitionService   TransitionServiceInterface
	TrackerService      TrackerServiceInterface
	SeedService         SeedServiceInterface
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateProductDatabase(DB)
	return Domain{
		ProjectService: ProjectService{
			ProjectRepository: ProjectRepository{
				DB: DB,
			},
		},
		IssueService: IssueService{
			IssueRepository: IssueRepository{
				DB: DB,
			},
		},
		FilterService: FilterService{
			FilterRepository: FilterRepository{
				DB: DB,
			},
		},
		BoardService: BoardService{
			BoardRepository: BoardRepository{
				DB: DB,
			},
		},
		DashboardService: DashboardService{
			DashboardRepository: DashboardRepository{
				DB: DB,
			},
		},
		UserService: UserService{
			UserRepository: UserRepository{
				DB: DB,
			},
		},
		UserSessionService: UserSessionService{
			UserRepository: UserRepository{
				DB: DB,
			},
		},
		UserRegisterService: UserRegisterService{
			UserRepository: UserRepository{
				DB: DB,
			},
		},
		StatusService: StatusService{
			StatusRepository: StatusRepository{
				DB: DB,
			},
		},
		TransitionService: TransitionService{
			TransitionRepository: TransitionRepository{
				DB: DB,
			},
		},
		TrackerService: TrackerService{
			TrackerRepository: TrackerRepository{
				DB: DB,
			},
		},
		SeedService: SeedService{
			StatusRepository: StatusRepository{
				DB: DB,
			},
		},
	}
}
