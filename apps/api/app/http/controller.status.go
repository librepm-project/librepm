package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
)

type StatusControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type StatusController struct {
	StatusService domain.StatusServiceInterface
}

func (c StatusController) Index(w http.ResponseWriter, r *http.Request) {
	statuses, err := c.StatusService.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponseMany(*statuses))
}

func (c StatusController) Show(w http.ResponseWriter, r *http.Request) {
	var status_id, _ = http_utils.GetParamUUID(r, "status_id")
	status, err := c.StatusService.Show(status_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponse(*status))
}

func (c StatusController) Create(w http.ResponseWriter, r *http.Request) {
	var status_request StatusRequest
	json.NewDecoder(r.Body).Decode(&status_request)
	status := StatusSerializer{}.RequestToModel(status_request)
	err := c.StatusService.Create(&status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, StatusSerializer{}.ModelToResponse(status))
}

func (c StatusController) Update(w http.ResponseWriter, r *http.Request) {
	status_id, _ := http_utils.GetParamUUID(r, "status_id")
	var status_request StatusRequest
	json.NewDecoder(r.Body).Decode(&status_request)
	status := StatusSerializer{}.RequestToModel(status_request)
	err := c.StatusService.Update(status_id, &status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updated, err := c.StatusService.Show(status_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponse(*updated))
}

func (c StatusController) Destroy(w http.ResponseWriter, r *http.Request) {
	status_id, _ := http_utils.GetParamUUID(r, "status_id")
	err := c.StatusService.Destroy(status_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
