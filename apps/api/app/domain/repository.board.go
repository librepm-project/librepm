package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardRepositoryInterface interface {
	All() (*[]BoardModel, error)
	FindByID(board_id uuid.UUID) (*BoardModel, error)
	Create(board *BoardModel) error
	Update(board_id uuid.UUID, board *BoardModel) error
	Destroy(board_id uuid.UUID) error
}

type BoardRepository struct {
	DB *gorm.DB
}

func (r BoardRepository) All() (*[]BoardModel, error) {
	var boards []BoardModel
	var err error
	query := r.DB.Select("board.*")

	if err := query.Find(&boards).Error; err != nil {
		fmt.Println(err)
	}
	return &boards, err
}

func (r BoardRepository) FindByID(board_id uuid.UUID) (*BoardModel, error) {
	var board BoardModel
	query := r.DB.Model(BoardModel{ID: board_id}).Scan(&board)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &board, query.Error
}

func (r BoardRepository) Create(board *BoardModel) error {
	query := r.DB.Create(&board)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r BoardRepository) Update(board_id uuid.UUID, board *BoardModel) error {
	query := r.DB.Model(BoardModel{}).Where("id", board_id).Updates(&board)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r BoardRepository) Destroy(board_id uuid.UUID) error {
	query := r.DB.Model(BoardModel{}).Delete(BoardModel{}, board_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
