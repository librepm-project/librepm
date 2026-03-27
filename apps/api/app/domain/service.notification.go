package domain

import "github.com/google/uuid"

type NotificationPusher interface {
	PushToUser(userID uuid.UUID, notification *NotificationModel)
}

type NotificationServiceInterface interface {
	AllUnread(userID uuid.UUID) (*[]NotificationModel, error)
	MarkAsRead(id uuid.UUID) error
	Create(notification *NotificationModel) error
	NotifyIssueParticipants(issue *IssueModel, notifType string, actorID uuid.UUID)
	SetPusher(p NotificationPusher)
}

type NotificationService struct {
	NotificationRepository NotificationRepositoryInterface
	Pusher                 NotificationPusher
}

func (s *NotificationService) SetPusher(p NotificationPusher) {
	s.Pusher = p
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

func (s NotificationService) NotifyIssueParticipants(issue *IssueModel, notifType string, actorID uuid.UUID) {
	entityType := EntityTypeIssue
	recipients := map[uuid.UUID]bool{}
	if issue.ReporterUserID != nil && *issue.ReporterUserID != actorID {
		recipients[*issue.ReporterUserID] = true
	}
	if issue.AssignedUserID != nil && *issue.AssignedUserID != actorID {
		recipients[*issue.AssignedUserID] = true
	}
	for recipientID := range recipients {
		n := &NotificationModel{
			UserID:     recipientID,
			Type:       notifType,
			EntityType: &entityType,
			EntityID:   &issue.ID,
		}
		if err := s.NotificationRepository.Create(n); err == nil && s.Pusher != nil {
			s.Pusher.PushToUser(recipientID, n)
		}
	}
}
