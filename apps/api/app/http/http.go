package http

import (
	"apps/api/app/domain"
	"fmt"
	"libs/http_utils"
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
			IssueService: domain.IssueService,
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
		TrackerController: TrackerController{
			TrackerService: domain.TrackerService,
		},
		TransitionController: TransitionController{
			TransitionService: domain.TransitionService,
		},
	}.Init()

	fmt.Println("API listen on port 80")
	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router:                         router,
		AuthorizationMiddlewareContext: authorization_middleware_context,
	})
}
