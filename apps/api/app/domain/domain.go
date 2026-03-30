package domain

import (
	"gorm.io/gorm"
)

// Entity type constants
const (
	EntityTypeIssue   = "issue"
	EntityTypeProject = "project"
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
	OnboardService          OnboardServiceInterface
	StatusService           StatusServiceInterface
	PriorityService         PriorityServiceInterface
	TransitionService       TransitionServiceInterface
	TrackerService          TrackerServiceInterface
	AttachmentService       AttachmentServiceInterface
	IssueAuditLogService    IssueAuditLogServiceInterface
	CommentService          CommentServiceInterface
	SettingService          SettingServiceInterface
	NotificationService     NotificationServiceInterface
	ReleaseService          ReleaseServiceInterface
	ProjectReleaseService   ProjectReleaseServiceInterface
	ProjectReleaseIssueService ProjectReleaseIssueServiceInterface

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
	PriorityRepository             PriorityRepositoryInterface
	TransitionRepository           TransitionRepositoryInterface
	TrackerRepository              TrackerRepositoryInterface
	ProjectTrackerRepository       ProjectTrackerRepositoryInterface
	FilterConditionRepository      FilterConditionRepositoryInterface
	AttachmentRepository           AttachmentRepositoryInterface
	IssueAuditLogRepository        IssueAuditLogRepositoryInterface
	CommentRepository              CommentRepositoryInterface
	SettingRepository              SettingRepositoryInterface
	NotificationRepository         NotificationRepositoryInterface
	ReleaseRepository              ReleaseRepositoryInterface
	ProjectReleaseRepository       ProjectReleaseRepositoryInterface
	ProjectReleaseIssueRepository  ProjectReleaseIssueRepositoryInterface
}

func NewDomain(DB *gorm.DB) Domain {
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
	priorityRepository := PriorityRepository{DB: DB}
	transitionRepository := TransitionRepository{DB: DB}
	trackerRepository := TrackerRepository{DB: DB}
	projectTrackerRepository := ProjectTrackerRepository{DB: DB}
	filterConditionRepository := FilterConditionRepository{DB: DB}
	attachmentRepository := AttachmentRepository{DB: DB}
	issueAuditLogRepository := IssueAuditLogRepository{DB: DB}
	commentRepository := CommentRepository{DB: DB}
	settingRepository := SettingRepository{DB: DB}
	settingRepository.Seed()
	notificationRepository := NotificationRepository{DB: DB}
	releaseRepository := ReleaseRepository{DB: DB}
	projectReleaseRepository := ProjectReleaseRepository{DB: DB}
	projectReleaseIssueRepository := ProjectReleaseIssueRepository{DB: DB}

	return Domain{
		DB: DB,

		ProjectService: ProjectService{
			ProjectRepository:              projectRepository,
			ProjectIssuePropertyRepository: projectIssuePropertyRepository,
			ProjectTrackerRepository:       projectTrackerRepository,
			ProjectTrackerStatusRepository: projectTrackerStatusRepository,
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
		OnboardService: OnboardService{
			SettingRepository: settingRepository,
			UserRepository:    userRepository,
		},
		StatusService: StatusService{
			StatusRepository: statusRepository,
		},
		PriorityService: PriorityService{
			PriorityRepository: priorityRepository,
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
		IssueAuditLogService: IssueAuditLogService{
			IssueAuditLogRepository: issueAuditLogRepository,
		},
		CommentService: CommentService{
			CommentRepository: commentRepository,
		},
		SettingService: SettingService{
			SettingRepository: settingRepository,
		},
		NotificationService: &NotificationService{
			NotificationRepository: notificationRepository,
		},
		ReleaseService: ReleaseService{
			ReleaseRepository: releaseRepository,
		},
		ProjectReleaseService: ProjectReleaseService{
			ProjectReleaseRepository: projectReleaseRepository,
		},
		ProjectReleaseIssueService: ProjectReleaseIssueService{
			ProjectReleaseIssueRepository: projectReleaseIssueRepository,
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
		PriorityRepository:             priorityRepository,
		TransitionRepository:           transitionRepository,
		TrackerRepository:              trackerRepository,
		ProjectTrackerRepository:       projectTrackerRepository,
		FilterConditionRepository:      filterConditionRepository,
		ProjectTrackerStatusRepository: projectTrackerStatusRepository,
		ProjectIssuePropertyRepository: projectIssuePropertyRepository,
		AttachmentRepository:           attachmentRepository,
		IssueAuditLogRepository:        issueAuditLogRepository,
		CommentRepository:              commentRepository,
		SettingRepository:              settingRepository,
		NotificationRepository:         notificationRepository,
		ReleaseRepository:              releaseRepository,
		ProjectReleaseRepository:       projectReleaseRepository,
		ProjectReleaseIssueRepository:  projectReleaseIssueRepository,
	}
}
