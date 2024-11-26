package domain

import (
	"gorm.io/gorm"
)

func MigrateProductDatabase(db *gorm.DB) {
	db.AutoMigrate(
		BoardModel{},
		BoardColumnIssueStatusModel{},
		BoardColumnModel{},
		DashboardModel{},
		FilterConditionModel{},
		FilterModel{},
		IssueModel{},
		IssueStatusModel{},
		IssueTransitionModel{},
		IssueTypeModel{},
		ProjectIssueTypeModel{},
		ProjectIssueTypeStatusModel{},
		ProjectIssueTypeTransitionModel{},
		ProjectModel{},
		ProjectUserModel{},
		UserModel{},
	)
}
