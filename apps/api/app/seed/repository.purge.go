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

		&domain.IssueModel{},

		&domain.ProjectTrackerStatusModel{},
		&domain.ProjectTrackerModel{},
		&domain.ProjectModel{},

		&domain.DashboardWidgetModel{},
		&domain.DashboardModel{},

		&domain.BoardColumnModel{},
		&domain.BoardModel{},

		&domain.FilterConditionModel{},
		&domain.FilterModel{},

		&domain.StatusModel{},
		&domain.TrackerModel{},

		&domain.NotificationModel{},
		&domain.UserModel{},
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
