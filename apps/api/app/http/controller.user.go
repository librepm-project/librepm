package http

import (
	"net/http"

	"apps/api/app/domain"
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
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, UserSerializer{}.ModelToResponseMany(*users))
}

func (c UserController) Show(w http.ResponseWriter, r *http.Request) {
	user_id, err := GetParamUUID(r, "user_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	user, err := c.UserService.Show(user_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, UserSerializer{}.ModelToResponse(*user))
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user_request UserRequest
	if err := DecodeJSON(r, &user_request); err != nil {
		RespondBadRequest(w)
		return
	}
	user := UserSerializer{}.RequestToModel(user_request)
	err := c.UserService.Create(&user)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, UserSerializer{}.ModelToResponse(user))
}

func (c UserController) Update(w http.ResponseWriter, r *http.Request) {
	user_id, err := GetParamUUID(r, "user_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var user_request UserRequest
	if err := DecodeJSON(r, &user_request); err != nil {
		RespondBadRequest(w)
		return
	}
	user := UserSerializer{}.RequestToModel(user_request)
	err = c.UserService.Update(user_id, &user)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	updated, err := c.UserService.Show(user_id)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, UserSerializer{}.ModelToResponse(*updated))
}

func (c UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	user_id, err := GetParamUUID(r, "user_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	err = c.UserService.Destroy(user_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
