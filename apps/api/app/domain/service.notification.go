package domain

import "github.com/google/uuid"

type NotificationServiceInterface interface {
	AllUnread(userID uuid.UUID) (*[]NotificationModel, error)
	MarkAsRead(id uuid.UUID) error
	Create(notification *NotificationModel) error
	NotifyIssueParticipants(issue *IssueModel, notifType string, actorID uuid.UUID)
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
		s.NotificationRepository.Create(&NotificationModel{
			UserID:     recipientID,
			Type:       notifType,
			EntityType: &entityType,
			EntityID:   &issue.ID,
		})
	}
}
