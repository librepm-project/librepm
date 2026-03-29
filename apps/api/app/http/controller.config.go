package http

import (
	"net/http"
	"os"
)

type ConfigResponse struct {
	RegisterAllowed bool `json:"registerAllowed"`
}

type ConfigController struct{}

func (c ConfigController) Show(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, ConfigResponse{
		RegisterAllowed: os.Getenv("REGISTER_ALLOWED") == "true",
	})
}
