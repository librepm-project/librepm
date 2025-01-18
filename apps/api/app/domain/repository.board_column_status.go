package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnStatusRepositoryInterface interface {
	All() (*[]BoardColumnModel, error)
	FindByID(board_column_status_id uuid.UUID) (*BoardColumnModel, error)
	Create(board_column_status *BoardColumnModel) error
	Update(board_column_status_id uuid.UUID, board_column_status *BoardColumnModel) error
	Destroy(board_column_status_id uuid.UUID) error
}

type BoardColumnStatusRepository struct {
	DB *gorm.DB
}

func (r BoardColumnStatusRepository) All() (*[]BoardColumnModel, error) {
	var board_column_statuses []BoardColumnModel
	var err error
	query := r.DB.Select("board_column_status.*")

	if err := query.Find(&board_column_statuses).Error; err != nil {
		fmt.Println(err)
	}
	return &board_column_statuses, err
}

func (r BoardColumnStatusRepository) FindByID(board_column_status_id uuid.UUID) (*BoardColumnModel, error) {
	var board_column_status BoardColumnModel
	query := r.DB.Model(BoardColumnModel{ID: board_column_status_id}).Scan(&board_column_status)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &board_column_status, query.Error
}

func (r BoardColumnStatusRepository) Create(board_column_status *BoardColumnModel) error {
	query := r.DB.Create(&board_column_status)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r BoardColumnStatusRepository) Update(board_column_status_id uuid.UUID, board_column_status *BoardColumnModel) error {
	query := r.DB.Model(BoardColumnModel{}).Where("id", board_column_status_id).Updates(&board_column_status)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r BoardColumnStatusRepository) Destroy(board_column_status_id uuid.UUID) error {
	query := r.DB.Model(BoardColumnModel{}).Delete(BoardColumnModel{}, board_column_status_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
