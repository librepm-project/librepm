package domain

import (
	"gorm.io/gorm"
)

func MigrateProductDatabase(db *gorm.DB) {
	db.AutoMigrate(
		BoardColumnIssueStatusModel{},
		BoardColumnModel{},
		BoardModel{},
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
