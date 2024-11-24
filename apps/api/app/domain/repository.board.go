package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardRepositoryInterface interface {
	All() *[]BoardModel
	FindByID(board_id uuid.UUID) *BoardModel
	Create(board *BoardModel)
	Update(board_id uuid.UUID, board *BoardModel)
	Destroy(board_id uuid.UUID)
}

type BoardRepository struct {
	DB *gorm.DB
}

func (r BoardRepository) All() *[]BoardModel {
	var boards []BoardModel
	query := r.DB.Select("board.*")

	if err := query.Find(&boards).Error; err != nil {
		fmt.Println(err)
	}
	return &boards
}

func (r BoardRepository) FindByID(board_id uuid.UUID) *BoardModel {
	var board BoardModel
	fmt.Println(board_id)
	err := r.DB.Model(BoardModel{ID: board_id}).Scan(&board)

	if err != nil {
		fmt.Println(err)
	}
	return &board
}

func (r BoardRepository) Create(board *BoardModel) {
	r.DB.Create(&board)
}

func (r BoardRepository) Update(board_id uuid.UUID, board *BoardModel) {
	r.DB.Model(BoardModel{}).Where("id", board_id).Updates(&board)
}

func (r BoardRepository) Destroy(board_id uuid.UUID) {
	r.DB.Model(BoardModel{}).Delete(BoardModel{}, board_id)
}
