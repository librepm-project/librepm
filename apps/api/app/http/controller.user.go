package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
)

type UserControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	UserService domain.UserServiceInterface
}

func (c UserController) Index(w http.ResponseWriter, r *http.Request) {
	users := c.UserService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.SerializeUsers(*users))
}

func (c UserController) Show(w http.ResponseWriter, r *http.Request) {
	var user_id, _ = http_utils.GetParamUUID(r, "user_id")
	user := c.UserService.Show(user_id)
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.SerializeUser(*user))
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.UserModel
	json.NewDecoder(r.Body).Decode(&user)
	c.UserService.Create(&user)
	http_utils.RespondWithJSON(w, http.StatusCreated, UserSerializer{}.SerializeUser(user))
}

func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	user_id, _ := http_utils.GetParamUUID(r, "user_id")
	var user domain.UserModel
	json.NewDecoder(r.Body).Decode(&user)
	c.UserService.Update(user_id, &user)
	w.WriteHeader(http.StatusOK)
}

func (c UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	user_id, _ := http_utils.GetParamUUID(r, "user_id")
	c.UserService.Destroy(user_id)
	w.WriteHeader(http.StatusNoContent)
}
