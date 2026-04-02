package domain

import (
	"fmt"

	"libs/password_utils"
)

type OnboardServiceInterface interface {
	Execute(siteTitle string, email string, password string, firstName string, lastName string) error
}

type OnboardService struct {
	SettingRepository  SettingRepositoryInterface
	UserRepository     UserRepositoryInterface
	UserRoleRepository UserRoleRepositoryInterface
}

func (s OnboardService) Execute(siteTitle, email, password, firstName, lastName string) error {
	setting, err := s.SettingRepository.FindByKey(SettingKeyOnboarded)
	if err == nil && setting.Value == "true" {
		return fmt.Errorf("already onboarded")
	}

	if err := s.SettingRepository.Update(SettingKeySiteTitle, siteTitle); err != nil {
		return err
	}

	passwordHash, _ := password_utils.HashPassword(password)
	user := UserModel{
		Email:        email,
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
	}
	if err := s.UserRepository.Create(&user); err != nil {
		return err
	}

	adminRole := "admin"
	_ = s.UserRoleRepository.Create(&UserRoleModel{
		UserID: user.ID,
		Role:   adminRole,
	})

	return s.SettingRepository.Update(SettingKeyOnboarded, "true")
}
