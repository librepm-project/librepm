package http

import (
	"apps/api/app/domain"
	"encoding/json"
	"net/http"

	"libs/http_utils"
)

type UserRegisterControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type UserRegisterController struct {
	UserRegisterService domain.UserRegisterServiceInterface
}

func (c UserRegisterController) Create(w http.ResponseWriter, r *http.Request) {
	var body UserRegisterRequest
	json.NewDecoder(r.Body).Decode(&body)
	user, err := c.UserRegisterService.Create(body.Email, body.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.ModelToResponse(*user))
}
