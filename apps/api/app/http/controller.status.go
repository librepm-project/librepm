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
	statuses := c.StatusService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponseMany(*statuses))
}

func (c StatusController) Show(w http.ResponseWriter, r *http.Request) {
	var status_id, _ = http_utils.GetParamUUID(r, "status_id")
	status := c.StatusService.Show(status_id)
	http_utils.RespondWithJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponse(*status))
}

func (c StatusController) Create(w http.ResponseWriter, r *http.Request) {
	var status_request StatusRequest
	json.NewDecoder(r.Body).Decode(&status_request)
	status := StatusSerializer{}.RequestToModel(status_request)
	c.StatusService.Create(&status)
	http_utils.RespondWithJSON(w, http.StatusCreated, StatusSerializer{}.ModelToResponse(status))
}

func (c StatusController) Update(w http.ResponseWriter, r *http.Request) {
	status_id, _ := http_utils.GetParamUUID(r, "status_id")
	var status_request StatusRequest
	json.NewDecoder(r.Body).Decode(&status_request)
	status := StatusSerializer{}.RequestToModel(status_request)
	c.StatusService.Update(status_id, &status)
	w.WriteHeader(http.StatusOK)
}

func (c StatusController) Destroy(w http.ResponseWriter, r *http.Request) {
	status_id, _ := http_utils.GetParamUUID(r, "status_id")
	c.StatusService.Destroy(status_id)
	w.WriteHeader(http.StatusNoContent)
}
