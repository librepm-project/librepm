package seed

import (
	"apps/api/app/domain"
	"libs/password_utils"
	"time"
)

func (s SeedService) createUser(items []UserData) error {
	var err error
	var user domain.UserModel
	for _, item := range items {
		hashedPassword, err := password_utils.HashPassword(item.PasswordHash)
		if err != nil {
			return err
		}
		user = domain.UserModel{
			Email:        item.Email,
			PasswordHash: hashedPassword,
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
