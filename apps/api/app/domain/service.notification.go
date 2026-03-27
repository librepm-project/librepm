package domain

import "github.com/google/uuid"

type NotificationServiceInterface interface {
	AllUnread(userID uuid.UUID) (*[]NotificationModel, error)
	MarkAsRead(id uuid.UUID) error
	Create(notification *NotificationModel) error
}

type NotificationService struct {
	NotificationRepository NotificationRepositoryInterface
}

func (s NotificationService) AllUnread(userID uuid.UUID) (*[]NotificationModel, error) {
	return s.NotificationRepository.AllUnreadByUser(userID)
}

func (s NotificationService) MarkAsRead(id uuid.UUID) error {
	return s.NotificationRepository.MarkAsRead(id)
}

func (s NotificationService) Create(notification *NotificationModel) error {
	return s.NotificationRepository.Create(notification)
}
