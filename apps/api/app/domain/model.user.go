package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key;"`
	Email        string    `gorm:"type:varchar(100);unique;not null;index;"`
	PasswordHash string    `gorm:"type:varchar(100); not null;"`
	FirstName    string    `gorm:"type:varchar(100);"`
	LastName     string    `gorm:"type:varchar(100);"`
	Phone        string    `gorm:"type:varchar(100);"`
	Language     string    `gorm:"type:varchar(3);"`
	Country      string    `gorm:"type:varchar(3);"`
	Blocked      bool
	ProjectUsers []ProjectUserModel `gorm:"foreignKey:UserID"`
	Filters      []FilterModel      `gorm:"foreignKey:UserID"`
	Dashboads    []DashboardModel   `gorm:"foreignKey:UserID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

func (user UserModel) TableName() string {
	return "user"
}

func (user *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
