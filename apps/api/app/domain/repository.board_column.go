package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnRepositoryInterface interface {
	All() *[]BoardColumnModel
	FindByID(board_column_id uuid.UUID) *BoardColumnModel
	Create(board_column *BoardColumnModel)
	Update(board_column_id uuid.UUID, board_column *BoardColumnModel)
	Destroy(board_column_id uuid.UUID)
}

type BoardColumnRepository struct {
	DB *gorm.DB
}

func (r BoardColumnRepository) All() *[]BoardColumnModel {
	var board_columns []BoardColumnModel
	query := r.DB.Select("board_column.*")

	if err := query.Find(&board_columns).Error; err != nil {
		fmt.Println(err)
	}
	return &board_columns
}

func (r BoardColumnRepository) FindByID(board_column_id uuid.UUID) *BoardColumnModel {
	var board_column BoardColumnModel
	fmt.Println(board_column_id)
	err := r.DB.Model(BoardColumnModel{ID: board_column_id}).Scan(&board_column)

	if err != nil {
		fmt.Println(err)
	}
	return &board_column
}

func (r BoardColumnRepository) Create(board_column *BoardColumnModel) {
	r.DB.Create(&board_column)
}

func (r BoardColumnRepository) Update(board_column_id uuid.UUID, board_column *BoardColumnModel) {
	r.DB.Model(BoardColumnModel{}).Where("id", board_column_id).Updates(&board_column)
}

func (r BoardColumnRepository) Destroy(board_column_id uuid.UUID) {
	r.DB.Model(BoardColumnModel{}).Delete(BoardColumnModel{}, board_column_id)
}
