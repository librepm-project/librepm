package http

import (
	"apps/api/app/domain"
	"libs/http_utils"
	"log/slog"
)

func StartHttpServer(domain domain.Domain) {
	authorization_middleware_context := http_utils.AuthorizationMiddlewareContext{
		Whitelist: []http_utils.Endpoint{
			{
				Method: "POST",
				Path:   "/user/session",
			},
			{
				Method: "POST",
				Path:   "/user/register",
			},
		},
	}

	router := Router{
		ProjectController: ProjectController{
			ProjectService: domain.ProjectService,
		},
		ProjectIssuePropertyController: ProjectIssuePropertyController{
			ProjectService: domain.ProjectService,
		},
		IssueController: IssueController{
			IssueService:         domain.IssueService,
			IssueAuditLogService: domain.IssueAuditLogService,
			NotificationService:  domain.NotificationService,
		},
		IssueWorklogController: IssueWorklogController{
			IssueWorklogService:  domain.IssueWorklogService,
			IssueAuditLogService: domain.IssueAuditLogService,
		},
		RelatedIssueController: RelatedIssueController{
			RelatedIssueService:  domain.RelatedIssueService,
			IssueAuditLogService: domain.IssueAuditLogService,
		},
		FilterController: FilterController{
			FilterService: domain.FilterService,
		},
		BoardController: BoardController{
			BoardService: domain.BoardService,
		},
		DashboardController: DashboardController{
			DashboardService: domain.DashboardService,
		},
		DashboardWidgetController: DashboardWidgetController{
			DashboardWidgetService: domain.DashboardWidgetService,
		},
		UserCurrentController: UserCurrentController{
			UserService: domain.UserService,
		},
		UserSessionController: UserSessionController{
			UserSessionService: domain.UserSessionService,
		},
		UserRegisterController: UserRegisterController{
			UserRegisterService: domain.UserRegisterService,
		},
		UserController: UserController{
			UserService: domain.UserService,
		},
		StatusController: StatusController{
			StatusService: domain.StatusService,
		},
		PriorityController: PriorityController{
			PriorityService: domain.PriorityService,
		},
		TrackerController: TrackerController{
			TrackerService: domain.TrackerService,
		},
		TransitionController: TransitionController{
			TransitionService: domain.TransitionService,
		},
		AttachmentController: AttachmentController{
			AttachmentService:    domain.AttachmentService,
			IssueAuditLogService: domain.IssueAuditLogService,
		},
		IssueAuditLogController: IssueAuditLogController{
			IssueAuditLogService: domain.IssueAuditLogService,
		},
		CommentController: CommentController{
			CommentService:      domain.CommentService,
			IssueService:        domain.IssueService,
			NotificationService: domain.NotificationService,
		},
		SettingController: SettingController{
			SettingService: domain.SettingService,
		},
		NotificationController: NotificationController{
			NotificationService: domain.NotificationService,
		},
	}.Init()

	slog.Info("API listen on port 80")
	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router:                         router,
		AuthorizationMiddlewareContext: authorization_middleware_context,
	})
}
