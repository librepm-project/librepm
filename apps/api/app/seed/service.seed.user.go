package seed

import (
	"apps/api/app/domain"
	"time"
)

func (s SeedService) createUser(items []UserData) error {
	var err error
	var user domain.UserModel
	for _, item := range items {
		user = domain.UserModel{
			Email:        item.Email,
			PasswordHash: item.PasswordHash,
			FirstName:    item.FirstName,
			LastName:     item.LastName,
			Phone:        item.Phone,
			Language:     item.Language,
			Country:      item.Country,
			Blocked:      item.Blocked,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		err = s.UserRepository.Create(&user)
	}
	return err
}
