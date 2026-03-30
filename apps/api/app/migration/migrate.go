package migration

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrationUser(),
		migrationStatus(),
		migrationTracker(),
		migrationPriority(),
		migrationBoard(),
		migrationSetting(),
		migrationFilter(),
		migrationDashboard(),
		migrationTransition(),
		migrationProject(),
		migrationProjectForeignKeys(),
		migrationBoardColumn(),
		migrationIssue(),
		migrationFilterCondition(),
		migrationFilterConditionCleanup(),
		migrationProjectTracker(),
		migrationProjectUser(),
		migrationDashboardWidget(),
		migrationBoardColumnStatus(),
		migrationProjectTrackerStatus(),
		migrationProjectTrackerTransition(),
		migrationWorklog(),
		migrationRelatedIssue(),
		migrationAttachment(),
		migrationComment(),
		migrationIssueAuditLog(),
		migrationNotification(),
		migrationRelease(),
		migrationProjectRelease(),
		migrationProjectReleaseIssue(),
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
