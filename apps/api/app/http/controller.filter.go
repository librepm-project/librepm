package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"libs/jwt_utils"
)

type FilterControllerInterface interface {
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type FilterController struct {
	FilterService domain.FilterServiceInterface
}

func (c FilterController) Index(w http.ResponseWriter, r *http.Request) {
	filters, _ := c.FilterService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.ModelToResponseMany(*filters))
}

func (c FilterController) Show(w http.ResponseWriter, r *http.Request) {
	var filter_id, _ = http_utils.GetParamUUID(r, "filter_id")
	filter, _ := c.FilterService.Show(filter_id)
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.ModelToResponse(*filter))
}

func (c FilterController) Create(w http.ResponseWriter, r *http.Request) {
	var filter_request FilterRequest
	json.NewDecoder(r.Body).Decode(&filter_request)
	filter := FilterSerializer{}.RequestToModel(filter_request)
	filter.UserID = jwt_utils.GetTokenInfoFromRequest(r).UserID
	c.FilterService.Create(&filter)
	http_utils.RespondWithJSON(w, http.StatusCreated, FilterSerializer{}.ModelToResponse(filter))
}

func (c FilterController) Update(w http.ResponseWriter, r *http.Request) {
	filter_id, _ := http_utils.GetParamUUID(r, "filter_id")
	var filter_request FilterRequest
	json.NewDecoder(r.Body).Decode(&filter_request)
	filter := FilterSerializer{}.RequestToModel(filter_request)
	c.FilterService.Update(filter_id, &filter)
	w.WriteHeader(http.StatusOK)
}

func (c FilterController) Destroy(w http.ResponseWriter, r *http.Request) {
	filter_id, _ := http_utils.GetParamUUID(r, "filter_id")
	c.FilterService.Destroy(filter_id)
	w.WriteHeader(http.StatusNoContent)
}
