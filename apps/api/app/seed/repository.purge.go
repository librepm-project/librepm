package seed

import (
	"apps/api/app/domain"
	"log"

	"gorm.io/gorm"
)

type PurgeRepositoryInterface interface {
	Purge()
}

type PurgeRepository struct {
	DB *gorm.DB
}

func (r PurgeRepository) Purge() {
	models := []interface{}{

		&domain.ProjectReleaseIssueModel{},
		&domain.ProjectReleaseModel{},
		&domain.ReleaseModel{},

		&domain.RelatedIssueModel{},
		&domain.IssueWorklogModel{},
		&domain.AttachmentModel{},
		&domain.CommentModel{},
		&domain.IssueAuditLogModel{},
		&domain.IssueModel{},

		&domain.ProjectTrackerTransitionModel{},
		&domain.ProjectTrackerStatusModel{},
		&domain.ProjectTrackerModel{},
		&domain.ProjectUserModel{},
		&domain.ProjectModel{},

		&domain.DashboardWidgetModel{},
		&domain.DashboardModel{},

		&domain.BoardColumnStatusModel{},
		&domain.BoardColumnModel{},
		&domain.BoardModel{},

		&domain.FilterConditionModel{},
		&domain.FilterModel{},

		&domain.TransitionModel{},
		&domain.StatusModel{},
		&domain.TrackerModel{},
		&domain.PriorityModel{},

		&domain.NotificationModel{},
		&domain.UserRoleModel{},
		&domain.UserModel{},
		&domain.SettingModel{},
	}
	if err := r.DB.Exec("SET FOREIGN_KEY_CHECKS = 0;").Error; err != nil {
		log.Printf("Failed to disable foreign key checks: %v", err)
		return
	}
	for _, m := range models {
		err := r.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(m).Error
		if err != nil {
			log.Printf("Purge error (%s): %v", m, err)
		}
	}
	if err := r.DB.Exec("SET FOREIGN_KEY_CHECKS = 1;").Error; err != nil {
		log.Printf("Failed to re-enable foreign key checks: %v", err)
	}
}
