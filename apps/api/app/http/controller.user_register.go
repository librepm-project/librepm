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

type UserRegisterCreateRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c UserRegisterController) Create(w http.ResponseWriter, r *http.Request) {
	var body UserRegisterCreateRequestBody
	json.NewDecoder(r.Body).Decode(&body)
	user := c.UserRegisterService.Create(body.Email, body.Password)
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.SerializeUser(*user))
}
