package http

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"libs/http_utils"
	"libs/jwt_utils"

	"github.com/google/uuid"
)

// GetUserIDFromRequest extracts user ID from JWT token
func GetUserIDFromRequest(r *http.Request) uuid.UUID {
	return jwt_utils.GetTokenInfoFromRequest(r).UserID
}

// RespondBadRequest sends a 400 Bad Request response
func RespondBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

// RespondNotFound sends a 404 Not Found response
func RespondNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

// RespondInternalError sends a 500 Internal Server Error response
func RespondInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

// RespondNoContent sends a 204 No Content response
func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// RespondForbidden sends a 403 Forbidden response
func RespondForbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}

// RespondJSON sends a JSON response with the given status code
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	http_utils.RespondWithJSON(w, statusCode, data)
}

// DecodeJSON decodes JSON from request body into the given struct
func DecodeJSON(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		slog.Error("failed to decode JSON", "error", err)
	}
	return err
}

// GetParam extracts a string parameter from the request URL
func GetParam(r *http.Request, key string) string {
	return http_utils.GetParam(r, key)
}

// GetParamUUID extracts a UUID parameter from the request URL
func GetParamUUID(r *http.Request, key string) (uuid.UUID, error) {
	return http_utils.GetParamUUID(r, key)
}
