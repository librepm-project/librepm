package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
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
	filters := c.FilterService.All()
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.SerializeFilters(*filters))
}

func (c FilterController) Show(w http.ResponseWriter, r *http.Request) {
	var filter_id, _ = http_utils.GetParamUUID(r, "filter_id")
	filter := c.FilterService.Show(filter_id)
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.SerializeFilter(*filter))
}

func (c FilterController) Create(w http.ResponseWriter, r *http.Request) {
	var filter domain.FilterModel
	json.NewDecoder(r.Body).Decode(&filter)
	c.FilterService.Create(&filter)
	http_utils.RespondWithJSON(w, http.StatusCreated, FilterSerializer{}.SerializeFilter(filter))
}

func (c FilterController) Update(w http.ResponseWriter, r *http.Request) {
	filter_id, _ := http_utils.GetParamUUID(r, "filter_id")
	var filter domain.FilterModel
	json.NewDecoder(r.Body).Decode(&filter)
	c.FilterService.Update(filter_id, &filter)
	w.WriteHeader(http.StatusOK)
}

func (c FilterController) Destroy(w http.ResponseWriter, r *http.Request) {
	filter_id, _ := http_utils.GetParamUUID(r, "filter_id")
	c.FilterService.Destroy(filter_id)
	w.WriteHeader(http.StatusNoContent)
}
