package domain

import (
	"libs/mysql_utils"

	"gorm.io/gorm"
)

type Domain struct {
	DB *gorm.DB

	ProjectService          ProjectServiceInterface
	IssueService            IssueServiceInterface
	IssueWorklogService     IssueWorklogServiceInterface
	RelatedIssueService     RelatedIssueServiceInterface
	FilterService           FilterServiceInterface
	BoardService            BoardServiceInterface
	DashboardService        DashboardServiceInterface
	DashboardWidgetService  DashboardWidgetServiceInterface
	UserService             UserServiceInterface
	UserSessionService      UserSessionServiceInterface
	UserRegisterService     UserRegisterServiceInterface
	StatusService           StatusServiceInterface
	TransitionService       TransitionServiceInterface
	TrackerService          TrackerServiceInterface
	AttachmentService       AttachmentServiceInterface

	ProjectRepository              ProjectRepositoryInterface
	ProjectTrackerStatusRepository ProjectTrackerStatusRepositoryInterface
	ProjectIssuePropertyRepository ProjectIssuePropertyRepositoryInterface
	IssueRepository                IssueRepositoryInterface
	IssueWorklogRepository         IssueWorklogRepositoryInterface
	RelatedIssueRepository         RelatedIssueRepositoryInterface
	FilterRepository               FilterRepositoryInterface
	BoardRepository                BoardRepositoryInterface
	BoardColumnRepository          BoardColumnRepositoryInterface
	BoardColumnStatusRepository    BoardColumnStatusRepositoryInterface
	DashboardRepository            DashboardRepositoryInterface
	DashboardWidgetRepository      DashboardWidgetRepositoryInterface
	UserRepository                 UserRepositoryInterface
	StatusRepository               StatusRepositoryInterface
	TransitionRepository           TransitionRepositoryInterface
	TrackerRepository              TrackerRepositoryInterface
	ProjectTrackerRepository       ProjectTrackerRepositoryInterface
	FilterConditionRepository      FilterConditionRepositoryInterface
	AttachmentRepository           AttachmentRepositoryInterface
}

func NewDomain() Domain {
	DB := mysql_utils.Init()
	MigrateProductDatabase(DB)

	projectRepository := ProjectRepository{DB: DB}
	projectTrackerStatusRepository := ProjectTrackerStatusRepository{DB: DB}
	projectIssuePropertyRepository := ProjectIssuePropertyRepository{DB: DB}
	issueRepository := IssueRepository{DB: DB}
	issueWorklogRepository := IssueWorklogRepository{DB: DB}
	relatedIssueRepository := RelatedIssueRepository{DB: DB}
	filterRepository := FilterRepository{DB: DB}
	boardRepository := BoardRepository{DB: DB}
	boardColumnRepository := BoardColumnRepository{DB: DB}
	boardColumnStatusRepository := BoardColumnStatusRepository{DB: DB}
	dashboardRepository := DashboardRepository{DB: DB}
	dashboardWidgetRepository := DashboardWidgetRepository{DB: DB}
	userRepository := UserRepository{DB: DB}
	statusRepository := StatusRepository{DB: DB}
	transitionRepository := TransitionRepository{DB: DB}
	trackerRepository := TrackerRepository{DB: DB}
	projectTrackerRepository := ProjectTrackerRepository{DB: DB}
	filterConditionRepository := FilterConditionRepository{DB: DB}
	attachmentRepository := AttachmentRepository{DB: DB}

	return Domain{
		DB: DB,

		ProjectService: ProjectService{
			ProjectRepository:              projectRepository,
			ProjectIssuePropertyRepository: projectIssuePropertyRepository,
		},
		IssueService: IssueService{
			IssueRepository:   issueRepository,
			ProjectRepository: projectRepository,
			FilterRepository:  filterRepository,
		},
		IssueWorklogService: IssueWorklogService{
			IssueWorklogRepository: issueWorklogRepository,
		},
		RelatedIssueService: RelatedIssueService{
			RelatedIssueRepository: relatedIssueRepository,
			IssueRepository:        issueRepository,
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
		DashboardWidgetService: DashboardWidgetService{
			DashboardWidgetRepository: dashboardWidgetRepository,
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
		AttachmentService: AttachmentService{
			AttachmentRepository: attachmentRepository,
		},

		ProjectRepository:              projectRepository,
		IssueRepository:                issueRepository,
		IssueWorklogRepository:         issueWorklogRepository,
		RelatedIssueRepository:         relatedIssueRepository,
		FilterRepository:               filterRepository,
		BoardRepository:                boardRepository,
		BoardColumnRepository:          boardColumnRepository,
		BoardColumnStatusRepository:    boardColumnStatusRepository,
		DashboardRepository:            dashboardRepository,
		DashboardWidgetRepository:      dashboardWidgetRepository,
		UserRepository:                 userRepository,
		StatusRepository:               statusRepository,
		TransitionRepository:           transitionRepository,
		TrackerRepository:              trackerRepository,
		ProjectTrackerRepository:       projectTrackerRepository,
		FilterConditionRepository:      filterConditionRepository,
		ProjectTrackerStatusRepository: projectTrackerStatusRepository,
		ProjectIssuePropertyRepository: projectIssuePropertyRepository,
		AttachmentRepository:           attachmentRepository,
	}
}
