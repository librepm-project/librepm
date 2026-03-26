package http

import (
	"apps/api/app/domain"
	"net/http"
)

type SettingControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type SettingController struct {
	SettingService domain.SettingServiceInterface
}

func (c SettingController) Index(w http.ResponseWriter, r *http.Request) {
	settings, err := c.SettingService.GetAll()
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, settings)
}

type updateSettingRequest struct {
	Value string `json:"value"`
}

func (c SettingController) Update(w http.ResponseWriter, r *http.Request) {
	key := GetParam(r, "key")
	var req updateSettingRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}

	if err := c.SettingService.Update(key, req.Value); err != nil {
		RespondInternalError(w)
		return
	}

	RespondNoContent(w)
}
