package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnRepositoryInterface interface {
	All() (*[]BoardColumnModel, error)
	FindByID(board_column_id uuid.UUID) (*BoardColumnModel, error)
	Create(board_column *BoardColumnModel) error
	Update(board_column_id uuid.UUID, board_column *BoardColumnModel) error
	Destroy(board_column_id uuid.UUID) error
}

type BoardColumnRepository struct {
	DB *gorm.DB
}

func (r BoardColumnRepository) All() (*[]BoardColumnModel, error) {
	var board_columns []BoardColumnModel
	var err error
	query := r.DB.Select("board_column.*")

	if err = query.Find(&board_columns).Error; err != nil {
		slog.Error("failed to fetch board columns", "error", err)
	}
	return &board_columns, err
}

func (r BoardColumnRepository) FindByID(board_column_id uuid.UUID) (*BoardColumnModel, error) {
	var board_column BoardColumnModel
	query := r.DB.First(&board_column, board_column_id)

	if query.Error != nil {
		slog.Error("failed to find board column by id", "error", query.Error)
	}
	return &board_column, query.Error
}

func (r BoardColumnRepository) Create(board_column *BoardColumnModel) error {
	query := r.DB.Create(&board_column)
	if query.Error != nil {
		slog.Error("failed to create board column", "error", query.Error)
	}
	return query.Error
}

func (r BoardColumnRepository) Update(board_column_id uuid.UUID, board_column *BoardColumnModel) error {
	query := r.DB.Model(BoardColumnModel{}).Where("id", board_column_id).Updates(&board_column)
	if query.Error != nil {
		slog.Error("failed to update board column", "error", query.Error)
	}
	return query.Error
}

func (r BoardColumnRepository) Destroy(board_column_id uuid.UUID) error {
	query := r.DB.Model(BoardColumnModel{}).Delete(BoardColumnModel{}, board_column_id)
	if query.Error != nil {
		slog.Error("failed to destroy board column", "error", query.Error)
	}
	return query.Error
}
