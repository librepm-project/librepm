package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
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
	transitions := c.TransitionService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, TransitionSerializer{}.ModelToResponseMany(*transitions))
}

func (c TransitionController) Show(w http.ResponseWriter, r *http.Request) {
	var transition_id, _ = http_utils.GetParamUUID(r, "transition_id")
	transition := c.TransitionService.Show(transition_id)
	http_utils.RespondWithJSON(w, http.StatusOK, TransitionSerializer{}.ModelToResponse(*transition))
}

func (c TransitionController) Create(w http.ResponseWriter, r *http.Request) {
	var transition_request TransitionRequest
	json.NewDecoder(r.Body).Decode(&transition_request)
	transition := TransitionSerializer{}.RequestToModel(transition_request)
	c.TransitionService.Create(&transition)
	http_utils.RespondWithJSON(w, http.StatusCreated, TransitionSerializer{}.ModelToResponse(transition))
}

func (c TransitionController) Update(w http.ResponseWriter, r *http.Request) {
	transition_id, _ := http_utils.GetParamUUID(r, "transition_id")
	var transition_request TransitionRequest
	json.NewDecoder(r.Body).Decode(&transition_request)
	transition := TransitionSerializer{}.RequestToModel(transition_request)
	c.TransitionService.Update(transition_id, &transition)
	w.WriteHeader(http.StatusOK)
}

func (c TransitionController) Destroy(w http.ResponseWriter, r *http.Request) {
	transition_id, _ := http_utils.GetParamUUID(r, "transition_id")
	c.TransitionService.Destroy(transition_id)
	w.WriteHeader(http.StatusNoContent)
}
