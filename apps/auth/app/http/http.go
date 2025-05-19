package http

import (
	"apps/auth/app/domain"
	"fmt"
	"libs/http_utils"
)

func StartHttpServer(domain domain.Domain) {
	authorization_middleware_context := http_utils.AuthorizationMiddlewareContext{
		Whitelist: []http_utils.Endpoint{},
	}

	router := Router{
		OAuth2Controller: OAuth2Controller{
			OAuth2Service: domain.OAuth2Service,
		},
	}.Init()

	fmt.Println("API listen on port 80")
	http_utils.StartGenericHTTPServer(http_utils.StartGenericHTTPServerContext{
		Router:                         router,
		AuthorizationMiddlewareContext: authorization_middleware_context,
	})
}
