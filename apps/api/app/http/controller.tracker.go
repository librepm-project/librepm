package http

import (
	"net/http"

	"apps/api/app/domain"
)

type TrackerControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type TrackerController struct {
	TrackerService domain.TrackerServiceInterface
}

func (c TrackerController) Index(w http.ResponseWriter, r *http.Request) {
	trackers, err := c.TrackerService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, TrackerSerializer{}.ModelToResponseMany(*trackers))
}

func (c TrackerController) Show(w http.ResponseWriter, r *http.Request) {
	tracker_id, err := GetParamUUID(r, "tracker_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	tracker, err := c.TrackerService.Show(tracker_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, TrackerSerializer{}.ModelToResponse(*tracker))
}

func (c TrackerController) Create(w http.ResponseWriter, r *http.Request) {
	var tracker_request TrackerRequest
	if err := DecodeJSON(r, &tracker_request); err != nil {
		RespondBadRequest(w)
		return
	}
	tracker := TrackerSerializer{}.RequestToModel(tracker_request)
	err := c.TrackerService.Create(&tracker)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusCreated, TrackerSerializer{}.ModelToResponse(tracker))
}

func (c TrackerController) Update(w http.ResponseWriter, r *http.Request) {
	tracker_id, err := GetParamUUID(r, "tracker_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	var tracker_request TrackerRequest
	if err := DecodeJSON(r, &tracker_request); err != nil {
		RespondBadRequest(w)
		return
	}
	tracker := TrackerSerializer{}.RequestToModel(tracker_request)
	err = c.TrackerService.Update(tracker_id, &tracker)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	updated, err := c.TrackerService.Show(tracker_id)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, TrackerSerializer{}.ModelToResponse(*updated))
}

func (c TrackerController) Destroy(w http.ResponseWriter, r *http.Request) {
	tracker_id, err := GetParamUUID(r, "tracker_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	err = c.TrackerService.Destroy(tracker_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
