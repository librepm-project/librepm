package http

import (
	"net/http"

	"apps/api/app/domain"
)

type PriorityControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type PriorityController struct {
	PriorityService domain.PriorityServiceInterface
}

func (c PriorityController) Index(w http.ResponseWriter, r *http.Request) {
	priorities, err := c.PriorityService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, PrioritySerializer{}.ModelToResponseMany(*priorities))
}

func (c PriorityController) Show(w http.ResponseWriter, r *http.Request) {
	priority_id, err := GetParamUUID(r, "priority_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	priority, err := c.PriorityService.Show(priority_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, PrioritySerializer{}.ModelToResponse(*priority))
}

func (c PriorityController) Create(w http.ResponseWriter, r *http.Request) {
	var req PriorityRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}
	priority := PrioritySerializer{}.RequestToModel(req)
	if err := c.PriorityService.Create(&priority); err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, PrioritySerializer{}.ModelToResponse(priority))
}

func (c PriorityController) Update(w http.ResponseWriter, r *http.Request) {
	priority_id, err := GetParamUUID(r, "priority_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var req PriorityRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}
	priority := PrioritySerializer{}.RequestToModel(req)
	if err := c.PriorityService.Update(priority_id, &priority); err != nil {
		RespondBadRequest(w)
		return
	}
	updated, err := c.PriorityService.Show(priority_id)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, PrioritySerializer{}.ModelToResponse(*updated))
}

func (c PriorityController) Destroy(w http.ResponseWriter, r *http.Request) {
	priority_id, err := GetParamUUID(r, "priority_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	if err := c.PriorityService.Destroy(priority_id); err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
