package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/stretchr/testify/assert"
)

func TestSettingService_GetAll_ReturnsSettings(t *testing.T) {
	expected := &[]domain.SettingModel{
		{Key: domain.SettingKeySiteTitle, Value: "LibrePM"},
	}
	repo := &mockrepo.MockSettingRepository{
		AllFunc: func() (*[]domain.SettingModel, error) { return expected, nil },
	}
	svc := domain.SettingService{SettingRepository: repo}

	result, err := svc.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSettingService_GetAll_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockSettingRepository{
		AllFunc: func() (*[]domain.SettingModel, error) { return nil, repoErr },
	}
	svc := domain.SettingService{SettingRepository: repo}

	_, err := svc.GetAll()

	assert.ErrorIs(t, err, repoErr)
}

func TestSettingService_GetByKey_ReturnsSetting(t *testing.T) {
	expected := &domain.SettingModel{Key: domain.SettingKeySiteTitle, Value: "MyPM"}
	repo := &mockrepo.MockSettingRepository{
		FindByKeyFunc: func(key string) (*domain.SettingModel, error) {
			assert.Equal(t, domain.SettingKeySiteTitle, key)
			return expected, nil
		},
	}
	svc := domain.SettingService{SettingRepository: repo}

	result, err := svc.GetByKey(domain.SettingKeySiteTitle)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSettingService_Update_DelegatesKeyAndValue(t *testing.T) {
	var updatedKey, updatedValue string
	repo := &mockrepo.MockSettingRepository{
		UpdateFunc: func(key string, value string) error {
			updatedKey = key
			updatedValue = value
			return nil
		},
	}
	svc := domain.SettingService{SettingRepository: repo}

	err := svc.Update(domain.SettingKeySiteTitle, "New Title")

	assert.NoError(t, err)
	assert.Equal(t, domain.SettingKeySiteTitle, updatedKey)
	assert.Equal(t, "New Title", updatedValue)
}
