package domain

import (
	"libs/mysql_utils"

	"gorm.io/gorm"
)

type Domain struct {
	DB *gorm.DB

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

	ProjectRepository              ProjectRepositoryInterface
	ProjectTrackerStatusRepository ProjectTrackerStatusRepositoryInterface
	ProjectIssuePropertyRepository ProjectIssuePropertyRepositoryInterface
	IssueRepository                IssueRepositoryInterface
	FilterRepository               FilterRepositoryInterface
	BoardRepository                BoardRepositoryInterface
	DashboardRepository            DashboardRepositoryInterface
	UserRepository                 UserRepositoryInterface
	StatusRepository               StatusRepositoryInterface
	TransitionRepository           TransitionRepositoryInterface
	TrackerRepository              TrackerRepositoryInterface
	ProjectTrackerRepository       ProjectTrackerRepositoryInterface
	FilterConditionRepository      FilterConditionRepositoryInterface
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateProductDatabase(DB)

	projectRepository := ProjectRepository{DB: DB}
	projectTrackerStatusRepository := ProjectTrackerStatusRepository{DB: DB}
	projectIssuePropertyRepository := ProjectIssuePropertyRepository{DB: DB}
	issueRepository := IssueRepository{DB: DB}
	filterRepository := FilterRepository{DB: DB}
	boardRepository := BoardRepository{DB: DB}
	dashboardRepository := DashboardRepository{DB: DB}
	userRepository := UserRepository{DB: DB}
	statusRepository := StatusRepository{DB: DB}
	transitionRepository := TransitionRepository{DB: DB}
	trackerRepository := TrackerRepository{DB: DB}
	projectTrackerRepository := ProjectTrackerRepository{DB: DB}
	filterConditionRepository := FilterConditionRepository{DB: DB}

	return Domain{
		DB: DB,

		ProjectService: ProjectService{
			ProjectRepository:              projectRepository,
			ProjectIssuePropertyRepository: projectIssuePropertyRepository,
		},
		IssueService: IssueService{
			IssueRepository:   issueRepository,
			ProjectRepository: projectRepository,
		},
		FilterService: FilterService{
			FilterRepository: filterRepository,
		},
		BoardService: BoardService{
			BoardRepository: boardRepository,
		},
		DashboardService: DashboardService{
			DashboardRepository: dashboardRepository,
		},
		UserService: UserService{
			UserRepository: userRepository,
		},
		UserSessionService: UserSessionService{
			UserRepository: userRepository,
		},
		UserRegisterService: UserRegisterService{
			UserRepository: userRepository,
		},
		StatusService: StatusService{
			StatusRepository: statusRepository,
		},
		TransitionService: TransitionService{
			TransitionRepository: transitionRepository,
		},
		TrackerService: TrackerService{
			TrackerRepository: trackerRepository,
		},

		ProjectRepository:              projectRepository,
		IssueRepository:                issueRepository,
		FilterRepository:               filterRepository,
		BoardRepository:                boardRepository,
		DashboardRepository:            dashboardRepository,
		UserRepository:                 userRepository,
		StatusRepository:               statusRepository,
		TransitionRepository:           transitionRepository,
		TrackerRepository:              trackerRepository,
		ProjectTrackerRepository:       projectTrackerRepository,
		FilterConditionRepository:      filterConditionRepository,
		ProjectTrackerStatusRepository: projectTrackerStatusRepository,
		ProjectIssuePropertyRepository: projectIssuePropertyRepository,
	}
}
