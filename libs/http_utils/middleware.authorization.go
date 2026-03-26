package http_utils

import (
	"fmt"
	"libs/jwt_utils"
	"net/http"
)

type Endpoint struct {
	Path   string
	Method string
}

type AuthorizationMiddlewareContext struct {
	Whitelist []Endpoint
}

func AuthorizationMiddleware(ctx AuthorizationMiddlewareContext) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
			for _, endpoint := range ctx.Whitelist {
				if endpoint.Path == r.URL.Path && endpoint.Method == r.Method {
					next.ServeHTTP(w, r)
					return
				}
			}
			token := jwt_utils.GetTokenString(r)
			if token == "" {
				unauthorized(w)
				return
			}
			token_info := jwt_utils.GetTokenInfo(token)
			if token_info == nil {
				unauthorized(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}
