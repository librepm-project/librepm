package http

import (
	"apps/api/app/domain"
	"net/http"
)

func RequirePermission(svc domain.PermissionServiceInterface, perm domain.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := GetUserIDFromRequest(r)
			if !svc.UserCan(userID, perm) {
				RespondForbidden(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
