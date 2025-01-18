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
	users, err := c.UserService.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.ModelToResponseMany(*users))
}

func (c UserController) Show(w http.ResponseWriter, r *http.Request) {
	var user_id, _ = http_utils.GetParamUUID(r, "user_id")
	user, err := c.UserService.Show(user_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, UserSerializer{}.ModelToResponse(*user))
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user_request UserRequest
	json.NewDecoder(r.Body).Decode(&user_request)
	user := UserSerializer{}.RequestToModel(user_request)
	err := c.UserService.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, UserSerializer{}.ModelToResponse(user))
}

func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	user_id, _ := http_utils.GetParamUUID(r, "user_id")
	var user_request UserRequest
	json.NewDecoder(r.Body).Decode(&user_request)
	user := UserSerializer{}.RequestToModel(user_request)
	err := c.UserService.Update(user_id, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	user_id, _ := http_utils.GetParamUUID(r, "user_id")
	err := c.UserService.Destroy(user_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
