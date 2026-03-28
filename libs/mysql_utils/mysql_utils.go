package mysql_utils

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConnectData struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func getDbConnectData() DbConnectData {
	return DbConnectData{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}
}

func Init() *gorm.DB {
	c := getDbConnectData()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)

	const maxAttempts = 30
	const retryInterval = 2 * time.Second

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil {
				if err = sqlDB.Ping(); err == nil {
					return db
				}
			}
		}
		slog.Info("waiting for database", "attempt", attempt, "max", maxAttempts)
		time.Sleep(retryInterval)
	}

	panic(fmt.Sprintf("failed to connect to database after %d attempts", maxAttempts))
}
