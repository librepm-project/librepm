package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepositoryInterface interface {
	AllUnreadByUser(userID uuid.UUID) (*[]NotificationModel, error)
	MarkAsRead(id uuid.UUID) error
	Create(notification *NotificationModel) error
}

type NotificationRepository struct {
	DB *gorm.DB
}

func (r NotificationRepository) AllUnreadByUser(userID uuid.UUID) (*[]NotificationModel, error) {
	var notifications []NotificationModel
	err := r.DB.Where("user_id = ? AND `read` = ?", userID, false).
		Order("created_at DESC").
		Find(&notifications).Error
	if err != nil {
		slog.Error("failed to fetch notifications", "error", err)
	}
	return &notifications, err
}

func (r NotificationRepository) MarkAsRead(id uuid.UUID) error {
	err := r.DB.Model(&NotificationModel{}).Where("id = ?", id).Update("read", true).Error
	if err != nil {
		slog.Error("failed to mark notification as read", "error", err)
	}
	return err
}

func (r NotificationRepository) Create(notification *NotificationModel) error {
	err := r.DB.Create(notification).Error
	if err != nil {
		slog.Error("failed to create notification", "error", err)
	}
	return err
}
