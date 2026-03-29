package http

import (
	"apps/api/app/domain"
	"libs/http_utils"
	"log/slog"
)

func StartHttpServer(d domain.Domain) {
	hub := NewHub()
	d.NotificationService.SetPusher(hub)

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
			{
				Method: "GET",
				Path:   "/ws",
			},
			{
				Method: "GET",
				Path:   "/setting",
			},
			{
				Method: "GET",
				Path:   "/config",
			},
		},
	}

	router := Router{
		ProjectController: ProjectController{
			ProjectService: d.ProjectService,
		},
		ProjectIssuePropertyController: ProjectIssuePropertyController{
			ProjectService: d.ProjectService,
		},
		IssueController: IssueController{
			IssueService:         d.IssueService,
			IssueAuditLogService: d.IssueAuditLogService,
			NotificationService:  d.NotificationService,
			Hub:                  hub,
		},
		IssueWorklogController: IssueWorklogController{
			IssueWorklogService:  d.IssueWorklogService,
			IssueAuditLogService: d.IssueAuditLogService,
		},
		RelatedIssueController: RelatedIssueController{
			RelatedIssueService:  d.RelatedIssueService,
			IssueAuditLogService: d.IssueAuditLogService,
		},
		FilterController: FilterController{
			FilterService: d.FilterService,
		},
		BoardController: BoardController{
			BoardService: d.BoardService,
		},
		DashboardController: DashboardController{
			DashboardService: d.DashboardService,
		},
		DashboardWidgetController: DashboardWidgetController{
			DashboardWidgetService: d.DashboardWidgetService,
		},
		UserCurrentController: UserCurrentController{
			UserService: d.UserService,
		},
		UserSessionController: UserSessionController{
			UserSessionService: d.UserSessionService,
		},
		UserRegisterController: UserRegisterController{
			UserRegisterService: d.UserRegisterService,
		},
		UserController: UserController{
			UserService: d.UserService,
		},
		StatusController: StatusController{
			StatusService: d.StatusService,
		},
		PriorityController: PriorityController{
			PriorityService: d.PriorityService,
		},
		TrackerController: TrackerController{
			TrackerService: d.TrackerService,
		},
		TransitionController: TransitionController{
			TransitionService: d.TransitionService,
		},
		AttachmentController: AttachmentController{
			AttachmentService:    d.AttachmentService,
			IssueAuditLogService: d.IssueAuditLogService,
		},
		IssueAuditLogController: IssueAuditLogController{
			IssueAuditLogService: d.IssueAuditLogService,
		},
		CommentController: CommentController{
			CommentService:      d.CommentService,
			IssueService:        d.IssueService,
			NotificationService: d.NotificationService,
		},
		SettingController: SettingController{
			SettingService: d.SettingService,
		},
		NotificationController: NotificationController{
			NotificationService: d.NotificationService,
		},
		WebSocketController: WebSocketController{
			Hub: hub,
		},
	}.Init()

	slog.Info("API listen on port 80")
	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router:                         router,
		AuthorizationMiddlewareContext: authorization_middleware_context,
	})
}
