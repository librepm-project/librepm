package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
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
	trackers := c.TrackerService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, TrackerSerializer{}.ModelToResponseMany(*trackers))
}

func (c TrackerController) Show(w http.ResponseWriter, r *http.Request) {
	var tracker_id, _ = http_utils.GetParamUUID(r, "tracker_id")
	tracker := c.TrackerService.Show(tracker_id)
	http_utils.RespondWithJSON(w, http.StatusOK, TrackerSerializer{}.ModelToResponse(*tracker))
}

func (c TrackerController) Create(w http.ResponseWriter, r *http.Request) {
	var tracker_request TrackerRequest
	json.NewDecoder(r.Body).Decode(&tracker_request)
	tracker := TrackerSerializer{}.RequestToModel(tracker_request)
	c.TrackerService.Create(&tracker)
	http_utils.RespondWithJSON(w, http.StatusCreated, TrackerSerializer{}.ModelToResponse(tracker))
}

func (c TrackerController) Update(w http.ResponseWriter, r *http.Request) {
	tracker_id, _ := http_utils.GetParamUUID(r, "tracker_id")
	var tracker_request TrackerRequest
	json.NewDecoder(r.Body).Decode(&tracker_request)
	tracker := TrackerSerializer{}.RequestToModel(tracker_request)
	c.TrackerService.Update(tracker_id, &tracker)
	w.WriteHeader(http.StatusOK)
}

func (c TrackerController) Destroy(w http.ResponseWriter, r *http.Request) {
	tracker_id, _ := http_utils.GetParamUUID(r, "tracker_id")
	c.TrackerService.Destroy(tracker_id)
	w.WriteHeader(http.StatusNoContent)
}
