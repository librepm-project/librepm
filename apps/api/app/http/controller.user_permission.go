package http

import (
	"apps/api/app/domain"
	"net/http"
)

type UserPermissionController struct {
	PermissionService domain.PermissionServiceInterface
}

func (c UserPermissionController) Index(w http.ResponseWriter, r *http.Request) {
	userID := GetUserIDFromRequest(r)
	permissions := c.PermissionService.UserPermissions(userID)
	RespondJSON(w, http.StatusOK, map[string][]string{"permissions": permissions})
}
