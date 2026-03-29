package http

import (
	"apps/api/app/domain"
	"encoding/json"
	"net/http"
	"os"

	"libs/http_utils"
)

type UserRegisterControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type UserRegisterController struct {
	UserRegisterService domain.UserRegisterServiceInterface
}

func (c UserRegisterController) Create(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("REGISTER_ALLOWED") != "true" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var body UserRegisterRequest
	json.NewDecoder(r.Body).Decode(&body)
	user, err := c.UserRegisterService.Create(body.Email, body.Password, body.FirstName, body.LastName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.ModelToResponse(*user))
}
