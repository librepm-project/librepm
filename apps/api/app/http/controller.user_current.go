package http

import (
	"apps/api/app/domain"
	"encoding/json"
	"net/http"

	"libs/http_utils"
	"libs/jwt_utils"
)

type UserCurrentControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type UserCurrentController struct {
	UserService domain.UserServiceInterface
}

func (c UserCurrentController) Show(w http.ResponseWriter, r *http.Request) {
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID
	user, _ := c.UserService.Show(user_id)
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.ModelToResponse(*user))
}

func (c UserCurrentController) Update(w http.ResponseWriter, r *http.Request) {
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID
	user := domain.UserModel{
		ID: user_id,
	}
	json.NewDecoder(r.Body).Decode(&user)
	c.UserService.Update(user_id, &user)
	w.WriteHeader(http.StatusOK)
}

func (c UserCurrentController) Destroy(w http.ResponseWriter, r *http.Request) {
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID
	c.UserService.Destroy(user_id)
	w.WriteHeader(http.StatusNoContent)
}
