package http

import (
	"apps/api/app/domain"
	"net/http"
)

type UserRoleController struct {
	UserRoleRepository domain.UserRoleRepositoryInterface
}

func (c UserRoleController) Index(w http.ResponseWriter, r *http.Request) {
	userID, err := GetParamUUID(r, "user_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	roles, err := c.UserRoleRepository.FindByUserID(userID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	result := make([]string, 0, len(roles))
	for _, r := range roles {
		result = append(result, r.Role)
	}
	RespondJSON(w, http.StatusOK, map[string][]string{"roles": result})
}

func (c UserRoleController) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := GetParamUUID(r, "user_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var body struct {
		Role string `json:"role"`
	}
	if err := DecodeJSON(r, &body); err != nil || body.Role == "" {
		RespondBadRequest(w)
		return
	}
	if _, ok := domain.Roles[body.Role]; !ok {
		RespondBadRequest(w)
		return
	}
	m := &domain.UserRoleModel{UserID: userID, Role: body.Role}
	if err := c.UserRoleRepository.Create(m); err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, map[string]string{"role": body.Role})
}

func (c UserRoleController) Destroy(w http.ResponseWriter, r *http.Request) {
	userID, err := GetParamUUID(r, "user_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	role := GetParam(r, "role")
	if role == "" {
		RespondBadRequest(w)
		return
	}
	if err := c.UserRoleRepository.Delete(userID, role); err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
