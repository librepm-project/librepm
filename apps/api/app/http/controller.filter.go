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
	Options(w http.ResponseWriter, r *http.Request)
}

type FilterController struct {
	FilterService domain.FilterServiceInterface
}

func (c FilterController) Index(w http.ResponseWriter, r *http.Request) {
	filters, err := c.FilterService.All()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.ModelToResponseMany(*filters))
}

func (c FilterController) Show(w http.ResponseWriter, r *http.Request) {
	var filter_id, _ = http_utils.GetParamUUID(r, "filter_id")
	filter, err := c.FilterService.Show(filter_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.ModelToResponse(*filter))
}

func (c FilterController) Create(w http.ResponseWriter, r *http.Request) {
	var filter_request FilterRequest
	json.NewDecoder(r.Body).Decode(&filter_request)
	filter := FilterSerializer{}.RequestToModel(filter_request)
	filter.UserID = jwt_utils.GetTokenInfoFromRequest(r).UserID
	err := c.FilterService.Create(&filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, FilterSerializer{}.ModelToResponse(filter))
}

func (c FilterController) Update(w http.ResponseWriter, r *http.Request) {
	filter_id, _ := http_utils.GetParamUUID(r, "filter_id")
	var filter_request FilterRequest
	json.NewDecoder(r.Body).Decode(&filter_request)
	filter := FilterSerializer{}.RequestToModel(filter_request)
	err := c.FilterService.Update(filter_id, &filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updated, err := c.FilterService.Show(filter_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, FilterSerializer{}.ModelToResponse(*updated))
}

func (c FilterController) Destroy(w http.ResponseWriter, r *http.Request) {
	filter_id, _ := http_utils.GetParamUUID(r, "filter_id")
	err := c.FilterService.Destroy(filter_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type filterConditionOperatorOption struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type filterConditionFieldOption struct {
	Key           string                           `json:"key"`
	Label         string                           `json:"label"`
	ValueType     string                           `json:"valueType"`
	ValueEndpoint string                           `json:"valueEndpoint,omitempty"`
	Operators     []filterConditionOperatorOption  `json:"operators"`
}

type filterConditionOptionsResponse struct {
	Fields []filterConditionFieldOption `json:"fields"`
}

func (c FilterController) Options(w http.ResponseWriter, r *http.Request) {
	eqNe := []filterConditionOperatorOption{
		{Key: "eq", Label: "is"},
		{Key: "ne", Label: "is not"},
	}
	options := filterConditionOptionsResponse{
		Fields: []filterConditionFieldOption{
			{
				Key:           "project_id",
				Label:         "Project",
				ValueType:     "select",
				ValueEndpoint: "/project",
				Operators:     eqNe,
			},
			{
				Key:           "tracker_id",
				Label:         "Tracker",
				ValueType:     "select",
				ValueEndpoint: "/tracker",
				Operators:     eqNe,
			},
			{
				Key:           "status_id",
				Label:         "Status",
				ValueType:     "select",
				ValueEndpoint: "/status",
				Operators:     eqNe,
			},
			{
				Key:       "summary",
				Label:     "Summary",
				ValueType: "text",
				Operators: []filterConditionOperatorOption{
					{Key: "contains", Label: "contains"},
					{Key: "eq", Label: "equals"},
				},
			},
		},
	}
	http_utils.RespondWithJSON(w, http.StatusOK, options)
}
