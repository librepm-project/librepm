package http

import (
	"apps/api/app/domain"
	"encoding/json"
	"net/http"

	"libs/http_utils"
)

type UserSessionControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type UserSessionController struct {
	UserSessionService domain.UserSessionServiceInterface
}

type UserSessionCreateRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c UserSessionController) Create(w http.ResponseWriter, r *http.Request) {
	var body UserSessionCreateRequestBody
	json.NewDecoder(r.Body).Decode(&body)
	session, err := c.UserSessionService.Create(body.Email, body.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.ToSessionResponse(*session))
}
