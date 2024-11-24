package http

import (
	"apps/api/app/domain"
	"libs/http_utils"
)

func StartHttpServer(domain domain.Domain) {
	authorization_middleware_context := http_utils.AuthorizationMiddlewareContext{
		Whitelist: []http_utils.Endpoint{},
	}

	router := Router{
		ProjectController: ProjectController{
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
	}.Init()

	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router:                         router,
		AuthorizationMiddlewareContext: authorization_middleware_context,
	})
}
