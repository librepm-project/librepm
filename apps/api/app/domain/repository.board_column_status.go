package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnStatusRepositoryInterface interface {
	All() *[]BoardColumnModel
	FindByID(board_column_status_id uuid.UUID) *BoardColumnModel
	Create(board_column_status *BoardColumnModel)
	Update(board_column_status_id uuid.UUID, board_column_status *BoardColumnModel)
	Destroy(board_column_status_id uuid.UUID)
}

type BoardColumnStatusRepository struct {
	DB *gorm.DB
}

func (r BoardColumnStatusRepository) All() *[]BoardColumnModel {
	var board_column_statuses []BoardColumnModel
	query := r.DB.Select("board_column_status.*")

	if err := query.Find(&board_column_statuses).Error; err != nil {
		fmt.Println(err)
	}
	return &board_column_statuses
}

func (r BoardColumnStatusRepository) FindByID(board_column_status_id uuid.UUID) *BoardColumnModel {
	var board_column_status BoardColumnModel
	fmt.Println(board_column_status_id)
	err := r.DB.Model(BoardColumnModel{ID: board_column_status_id}).Scan(&board_column_status)

	if err != nil {
		fmt.Println(err)
	}
	return &board_column_status
}

func (r BoardColumnStatusRepository) Create(board_column_status *BoardColumnModel) {
	r.DB.Create(&board_column_status)
}

func (r BoardColumnStatusRepository) Update(board_column_status_id uuid.UUID, board_column_status *BoardColumnModel) {
	r.DB.Model(BoardColumnModel{}).Where("id", board_column_status_id).Updates(&board_column_status)
}

func (r BoardColumnStatusRepository) Destroy(board_column_status_id uuid.UUID) {
	r.DB.Model(BoardColumnModel{}).Delete(BoardColumnModel{}, board_column_status_id)
}
