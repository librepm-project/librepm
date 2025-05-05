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
		&domain.ProjectTrackerModel{},
		&domain.ProjectModel{},

		&domain.DashboardModel{},

		&domain.BoardColumnModel{},
		&domain.BoardModel{},

		&domain.FilterConditionModel{},
		&domain.FilterModel{},

		&domain.UserModel{},

		&domain.StatusModel{},
		&domain.TrackerModel{},
	}
	for _, m := range models {
		err := r.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(m).Error
		if err != nil {
			log.Printf("Purge error (%s): %v", m, err)
		}
	}
}
