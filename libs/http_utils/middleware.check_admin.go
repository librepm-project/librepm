package http_utils

import (
	"libs/jwt_utils"
	"net/http"
)

func CheckAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if jwt_utils.GetTokenInfoFromRequest(r).Admin {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return

		}
	})
}
