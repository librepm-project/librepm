package domain

import (
	"time"
)

type UserData struct {
	Email        string `yaml:"email"`
	PasswordHash string `yaml:"password_hash"`
	FirstName    string `yaml:"first_name"`
	LastName     string `yaml:"last_name"`
	Phone        string `yaml:"phone"`
	Language     string `yaml:"language"`
	Country      string `yaml:"country"`
	Blocked      bool   `yaml:"blocked"`
}

func (s SeedService) createUser(items []UserData) error {
	var err error
	var user UserModel
	for _, item := range items {
		user = UserModel{
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
