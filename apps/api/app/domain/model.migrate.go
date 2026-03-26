package domain

import (
	"gorm.io/gorm"
)

func MigrateProductDatabase(db *gorm.DB) {
	db.AutoMigrate(
		BoardModel{},
		IssueWorklogModel{},
		BoardColumnStatusModel{},
		BoardColumnModel{},
		DashboardModel{},
		DashboardWidgetModel{},
		FilterConditionModel{},
		FilterModel{},
		IssueModel{},
		RelatedIssueModel{},
		StatusModel{},
		TransitionModel{},
		TrackerModel{},
		ProjectTrackerModel{},
		ProjectTrackerStatusModel{},
		ProjectTrackerTransitionModel{},
		ProjectModel{},
		ProjectUserModel{},
		UserModel{},
		AttachmentModel{},
	)

	// A régi filter_condition schema cleanup:
	// Az eredeti condition_filter_id composite unique index (condition + filter_id)
	// és a condition NOT NULL oszlop akadályozza az új Field/Op/Value inserteket.
	// Az indexet előbb kell ledobni, különben a DROP COLUMN Error 1072-t dob.
	db.Exec("ALTER TABLE `filter_condition` DROP INDEX IF EXISTS `condition_filter_id`")
	db.Exec("ALTER TABLE `filter_condition` DROP COLUMN IF EXISTS `condition`")
}
