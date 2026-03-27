package http

import (
	"apps/api/app/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"libs/http_utils"
	"libs/jwt_utils"
)

type NotificationControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	MarkRead(w http.ResponseWriter, r *http.Request)
}

type NotificationController struct {
	NotificationService domain.NotificationServiceInterface
}

func (c NotificationController) Index(w http.ResponseWriter, r *http.Request) {
	userID := jwt_utils.GetTokenInfoFromRequest(r).UserID
	notifications, err := c.NotificationService.AllUnread(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, NotificationSerializer{}.ModelToResponseMany(*notifications))
}

func (c NotificationController) MarkRead(w http.ResponseWriter, r *http.Request) {
	notificationID, err := uuid.Parse(chi.URLParam(r, "notification_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.NotificationService.MarkAsRead(notificationID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
