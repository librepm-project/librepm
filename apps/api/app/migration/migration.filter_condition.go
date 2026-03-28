package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"apps/api/app/domain"
)

func migrationFilterCondition() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240101000014",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.FilterConditionModel{})
		},
	}
}

func migrationFilterConditionCleanup() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240101000015",
		Migrate: func(tx *gorm.DB) error {
			// A régi filter_condition schema cleanup:
			// Az eredeti condition_filter_id composite unique index (condition + filter_id)
			// és a condition NOT NULL oszlop akadályozza az új Field/Op/Value inserteket.
			// Az indexet előbb kell ledobni, különben a DROP COLUMN Error 1072-t dob.
			tx.Exec("ALTER TABLE `filter_condition` DROP INDEX IF EXISTS `condition_filter_id`")
			tx.Exec("ALTER TABLE `filter_condition` DROP COLUMN IF EXISTS `condition`")
			return nil
		},
	}
}
