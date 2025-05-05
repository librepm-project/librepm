package domain

import (
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
		&ProjectTrackerModel{},
		&ProjectModel{},

		&DashboardModel{},

		&BoardColumnModel{},
		&BoardModel{},

		&FilterConditionModel{},
		&FilterModel{},

		&UserModel{},

		&StatusModel{},
		&TrackerModel{},
	}
	for _, m := range models {
		err := r.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(m).Error
		if err != nil {
			log.Printf("Purge error (%s): %v", m, err)
		}
	}
}
