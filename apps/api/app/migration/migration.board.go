package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"apps/api/app/domain"
)

func migrationBoard() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240101000005",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.BoardModel{})
		},
	}
}
