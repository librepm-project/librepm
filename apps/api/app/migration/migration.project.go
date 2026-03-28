package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"apps/api/app/domain"
)

func migrationProject() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240101000010",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&domain.ProjectModel{})
		},
	}
}

func migrationProjectForeignKeys() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240101000011",
		Migrate: func(tx *gorm.DB) error {
			// Foreign key constraints for project default values (added separately because
			// inline REFERENCES in GORM type tags is not valid MariaDB syntax):
			tx.Exec("ALTER TABLE `project` DROP FOREIGN KEY `fk_project_default_status`")
			tx.Exec("ALTER TABLE `project` ADD CONSTRAINT `fk_project_default_status` FOREIGN KEY (`default_status_id`) REFERENCES `status`(`id`) ON DELETE SET NULL")
			tx.Exec("ALTER TABLE `project` DROP FOREIGN KEY `fk_project_default_tracker`")
			tx.Exec("ALTER TABLE `project` ADD CONSTRAINT `fk_project_default_tracker` FOREIGN KEY (`default_tracker_id`) REFERENCES `tracker`(`id`) ON DELETE SET NULL")
			return nil
		},
	}
}
