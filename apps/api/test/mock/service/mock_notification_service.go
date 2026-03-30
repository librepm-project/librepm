package mocksvc

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type MockNotificationService struct {
	AllUnreadFunc              func(userID uuid.UUID) (*[]domain.NotificationModel, error)
	MarkAsReadFunc             func(id uuid.UUID) error
	CreateFunc                 func(notification *domain.NotificationModel) error
	NotifyIssueParticipantsFunc func(issue *domain.IssueModel, notifType string, actorID uuid.UUID)
	SetPusherFunc              func(p domain.NotificationPusher)
}

func (m *MockNotificationService) AllUnread(userID uuid.UUID) (*[]domain.NotificationModel, error) {
	if m.AllUnreadFunc != nil {
		return m.AllUnreadFunc(userID)
	}
	return &[]domain.NotificationModel{}, nil
}

func (m *MockNotificationService) MarkAsRead(id uuid.UUID) error {
	if m.MarkAsReadFunc != nil {
		return m.MarkAsReadFunc(id)
	}
	return nil
}

func (m *MockNotificationService) Create(notification *domain.NotificationModel) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(notification)
	}
	return nil
}

func (m *MockNotificationService) NotifyIssueParticipants(issue *domain.IssueModel, notifType string, actorID uuid.UUID) {
	if m.NotifyIssueParticipantsFunc != nil {
		m.NotifyIssueParticipantsFunc(issue, notifType, actorID)
	}
}

func (m *MockNotificationService) SetPusher(p domain.NotificationPusher) {
	if m.SetPusherFunc != nil {
		m.SetPusherFunc(p)
	}
}
