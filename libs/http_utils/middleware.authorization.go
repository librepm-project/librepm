package http_utils

import (
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
			for _, endpoint := range ctx.Whitelist {
				if endpoint.Path == r.URL.Path && endpoint.Method == r.Method {
					next.ServeHTTP(w, r)
					return
				}
			}
			token := r.Header.Get("Authorization")
			if token == "" {
				unauthorized(w)
				return
			}
			token_info := jwt_utils.GetTokenInfo((token))
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
