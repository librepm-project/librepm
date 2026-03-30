package mockrepo

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockNotificationRepository struct {
	AllUnreadByUserFunc func(userID uuid.UUID) (*[]domain.NotificationModel, error)
	MarkAsReadFunc      func(id uuid.UUID) error
	CreateFunc          func(notification *domain.NotificationModel) error
}

func (m *MockNotificationRepository) AllUnreadByUser(userID uuid.UUID) (*[]domain.NotificationModel, error) {
	if m.AllUnreadByUserFunc != nil {
		return m.AllUnreadByUserFunc(userID)
	}
	return &[]domain.NotificationModel{}, nil
}

func (m *MockNotificationRepository) MarkAsRead(id uuid.UUID) error {
	if m.MarkAsReadFunc != nil {
		return m.MarkAsReadFunc(id)
	}
	return nil
}

func (m *MockNotificationRepository) Create(notification *domain.NotificationModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(notification)
	}
	return nil
}
