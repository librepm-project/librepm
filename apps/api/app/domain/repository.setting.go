package domain

import (
	"encoding/json"
	"log/slog"

	"gorm.io/gorm"
)

type SettingRepositoryInterface interface {
	All() (*[]SettingModel, error)
	FindByKey(key string) (*SettingModel, error)
	Update(key string, value string) error
	Seed() error
}

type SettingRepository struct {
	DB *gorm.DB
}

func (r SettingRepository) All() (*[]SettingModel, error) {
	var settings []SettingModel
	if err := r.DB.Find(&settings).Error; err != nil {
		slog.Error("failed to fetch settings", "error", err)
		return nil, err
	}
	return &settings, nil
}

func (r SettingRepository) FindByKey(key string) (*SettingModel, error) {
	var setting SettingModel
	if err := r.DB.Where("`key` = ?", key).First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r SettingRepository) Update(key string, value string) error {
	if err := r.DB.Model(&SettingModel{}).Where("`key` = ?", key).Update("value", value).Error; err != nil {
		slog.Error("failed to update setting", "key", key, "error", err)
		return err
	}
	return nil
}

func (r SettingRepository) Seed() error {
	predefined := GetPredefinedSettings()
	for _, s := range predefined {
		var existing SettingModel
		if err := r.DB.Where("`key` = ?", s.Key).Limit(1).Find(&existing).Error; err != nil {
			slog.Error("failed to check setting during seed", "key", s.Key, "error", err)
			continue
		}
		if existing.Key == "" {
			if err := r.DB.Create(&s).Error; err != nil {
				slog.Error("failed to seed setting", "key", s.Key, "error", err)
			}
		} else {
			optionsJSON, err := json.Marshal(s.Options)
			if err != nil {
				slog.Error("failed to marshal setting options", "key", s.Key, "error", err)
				continue
			}
			if err := r.DB.Exec("UPDATE `setting` SET `options` = ? WHERE `key` = ?", string(optionsJSON), s.Key).Error; err != nil {
				slog.Error("failed to update setting options", "key", s.Key, "error", err)
			}
		}
	}
	return nil
}
