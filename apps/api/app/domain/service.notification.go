package domain

import "github.com/google/uuid"

type NotificationServiceInterface interface {
	AllUnread(userID uuid.UUID) (*[]NotificationModel, error)
	MarkAsRead(id uuid.UUID) error
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
