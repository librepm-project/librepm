package seed

import (
	"apps/api/app/domain"
	"time"
)

func (s SeedService) createNotification(items []NotificationData) error {
	var err error
	for _, item := range items {
		user, findErr := s.UserRepository.FindByEmail(item.UserEmail)
		if findErr != nil {
			return findErr
		}

		notification := domain.NotificationModel{
			UserID:    user.ID,
			Type:      item.Type,
			Read:      item.Read,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if item.EntityType != "" {
			notification.EntityType = &item.EntityType
		}

		err = s.NotificationRepository.Create(&notification)
	}
	return err
}
