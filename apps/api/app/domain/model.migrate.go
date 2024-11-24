package domain

import (
	"gorm.io/gorm"
)

func MigrateProductDatabase(db *gorm.DB) {
	db.AutoMigrate(
		ProjectModel{},
		IssueModel{},
	)
}
