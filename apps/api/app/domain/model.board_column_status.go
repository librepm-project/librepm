package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnStatusModel struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key;"`
	BoardColumnID uuid.UUID `gorm:"type:char(36) references board_column;not null;"`
	StatusID      uuid.UUID `gorm:"type:char(36) references status;not null;"`

	BoardColumn BoardColumnModel
	Status      StatusModel
}

func (board_column_status BoardColumnStatusModel) TableName() string {
	return "board_column_status"
}

func (board_column_status *BoardColumnStatusModel) BeforeCreate(tx *gorm.DB) (err error) {
	board_column_status.ID = uuid.New()
	return
}
