package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"apps/api/app/domain"
)

func migrationIssueAuditLog() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240101000026",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.IssueAuditLogModel{})
		},
	}
}
