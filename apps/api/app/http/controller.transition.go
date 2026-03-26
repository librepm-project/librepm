package http

import (
	"net/http"

	"apps/api/app/domain"
)

type TransitionControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type TransitionController struct {
	TransitionService domain.TransitionServiceInterface
}

func (c TransitionController) Index(w http.ResponseWriter, r *http.Request) {
	transitions, err := c.TransitionService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, TransitionSerializer{}.ModelToResponseMany(*transitions))
}

func (c TransitionController) Show(w http.ResponseWriter, r *http.Request) {
	transition_id, err := GetParamUUID(r, "transition_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	transition, err := c.TransitionService.Show(transition_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, TransitionSerializer{}.ModelToResponse(*transition))
}

func (c TransitionController) Create(w http.ResponseWriter, r *http.Request) {
	var transition_request TransitionRequest
	if err := DecodeJSON(r, &transition_request); err != nil {
		RespondBadRequest(w)
		return
	}
	transition := TransitionSerializer{}.RequestToModel(transition_request)
	err := c.TransitionService.Create(&transition)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, TransitionSerializer{}.ModelToResponse(transition))
}

func (c TransitionController) Update(w http.ResponseWriter, r *http.Request) {
	transition_id, err := GetParamUUID(r, "transition_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var transition_request TransitionRequest
	if err := DecodeJSON(r, &transition_request); err != nil {
		RespondBadRequest(w)
		return
	}
	transition := TransitionSerializer{}.RequestToModel(transition_request)
	err = c.TransitionService.Update(transition_id, &transition)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, TransitionSerializer{}.ModelToResponse(transition))
}

func (c TransitionController) Destroy(w http.ResponseWriter, r *http.Request) {
	transition_id, err := GetParamUUID(r, "transition_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	err = c.TransitionService.Destroy(transition_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
