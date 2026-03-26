package http

import (
	"net/http"

	"apps/api/app/domain"
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
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponseMany(*statuses))
}

func (c StatusController) Show(w http.ResponseWriter, r *http.Request) {
	status_id, err := GetParamUUID(r, "status_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	status, err := c.StatusService.Show(status_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponse(*status))
}

func (c StatusController) Create(w http.ResponseWriter, r *http.Request) {
	var status_request StatusRequest
	if err := DecodeJSON(r, &status_request); err != nil {
		RespondBadRequest(w)
		return
	}
	status := StatusSerializer{}.RequestToModel(status_request)
	err := c.StatusService.Create(&status)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, StatusSerializer{}.ModelToResponse(status))
}

func (c StatusController) Update(w http.ResponseWriter, r *http.Request) {
	status_id, err := GetParamUUID(r, "status_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var status_request StatusRequest
	if err := DecodeJSON(r, &status_request); err != nil {
		RespondBadRequest(w)
		return
	}
	status := StatusSerializer{}.RequestToModel(status_request)
	err = c.StatusService.Update(status_id, &status)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	updated, err := c.StatusService.Show(status_id)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, StatusSerializer{}.ModelToResponse(*updated))
}

func (c StatusController) Destroy(w http.ResponseWriter, r *http.Request) {
	status_id, err := GetParamUUID(r, "status_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	err = c.StatusService.Destroy(status_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
