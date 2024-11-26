package domain

import (
	"gorm.io/gorm"
)

func MigrateProductDatabase(db *gorm.DB) {
	db.AutoMigrate(
		BoardModel{},
		BoardColumnStatusModel{},
		BoardColumnModel{},
		DashboardModel{},
		FilterConditionModel{},
		FilterModel{},
		IssueModel{},
		StatusModel{},
		TransitionModel{},
		TrackerModel{},
		ProjectTrackerModel{},
		ProjectTrackerStatusModel{},
		ProjectTrackerTransitionModel{},
		ProjectModel{},
		ProjectUserModel{},
		UserModel{},
	)
}
