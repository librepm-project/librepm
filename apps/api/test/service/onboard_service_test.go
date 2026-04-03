package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/stretchr/testify/assert"
)

func TestOnboardService_Execute_ReturnsErrorWhenAlreadyOnboarded(t *testing.T) {
	settingRepo := &mockrepo.MockSettingRepository{
		FindByKeyFunc: func(key string) (*domain.SettingModel, error) {
			if key == domain.SettingKeyOnboarded {
				return &domain.SettingModel{Key: key, Value: "true"}, nil
			}
			return nil, errors.New("not found")
		},
	}
	svc := domain.OnboardService{
		SettingRepository: settingRepo,
		UserRepository:    &mockrepo.MockUserRepository{},
	}

	err := svc.Execute("LibrePM", "admin@example.com", "pass", "Admin", "User")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already onboarded")
}

func TestOnboardService_Execute_UpdatesTitleAndCreatesUser(t *testing.T) {
	var updatedKeys []string
	var createdUser *domain.UserModel

	settingRepo := &mockrepo.MockSettingRepository{
		FindByKeyFunc: func(key string) (*domain.SettingModel, error) {
			return nil, errors.New("not found")
		},
		UpdateFunc: func(key string, value string) error {
			updatedKeys = append(updatedKeys, key)
			return nil
		},
	}
	userRepo := &mockrepo.MockUserRepository{
		CreateFunc: func(user *domain.UserModel) error {
			createdUser = user
			return nil
		},
	}
	svc := domain.OnboardService{
		SettingRepository:  settingRepo,
		UserRepository:     userRepo,
		UserRoleRepository: &mockrepo.MockUserRoleRepository{},
	}

	err := svc.Execute("MyPM", "admin@example.com", "secure123", "Admin", "User")

	assert.NoError(t, err)
	assert.Contains(t, updatedKeys, domain.SettingKeySiteTitle)
	assert.Contains(t, updatedKeys, domain.SettingKeyOnboarded)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "admin@example.com", createdUser.Email)
	assert.Equal(t, "Admin", createdUser.FirstName)
}

func TestOnboardService_Execute_ReturnsErrorWhenUpdateTitleFails(t *testing.T) {
	updateErr := errors.New("update failed")
	settingRepo := &mockrepo.MockSettingRepository{
		FindByKeyFunc: func(key string) (*domain.SettingModel, error) {
			return nil, errors.New("not found")
		},
		UpdateFunc: func(key string, value string) error {
			return updateErr
		},
	}
	svc := domain.OnboardService{
		SettingRepository: settingRepo,
		UserRepository:    &mockrepo.MockUserRepository{},
	}

	err := svc.Execute("Title", "a@b.com", "pass", "A", "B")

	assert.ErrorIs(t, err, updateErr)
}

func TestOnboardService_Execute_ReturnsErrorWhenCreateUserFails(t *testing.T) {
	createErr := errors.New("create failed")
	settingRepo := &mockrepo.MockSettingRepository{
		FindByKeyFunc: func(key string) (*domain.SettingModel, error) {
			return nil, errors.New("not found")
		},
		UpdateFunc: func(key string, value string) error { return nil },
	}
	userRepo := &mockrepo.MockUserRepository{
		CreateFunc: func(user *domain.UserModel) error { return createErr },
	}
	svc := domain.OnboardService{
		SettingRepository: settingRepo,
		UserRepository:    userRepo,
	}

	err := svc.Execute("Title", "a@b.com", "pass", "A", "B")

	assert.ErrorIs(t, err, createErr)
}
