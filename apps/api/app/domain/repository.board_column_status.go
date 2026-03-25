package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardColumnStatusRepositoryInterface interface {
	All() (*[]BoardColumnStatusModel, error)
	FindByID(board_column_status_id uuid.UUID) (*BoardColumnStatusModel, error)
	Create(board_column_status *BoardColumnStatusModel) error
	Update(board_column_status_id uuid.UUID, board_column_status *BoardColumnStatusModel) error
	Destroy(board_column_status_id uuid.UUID) error
}

type BoardColumnStatusRepository struct {
	DB *gorm.DB
}

func (r BoardColumnStatusRepository) All() (*[]BoardColumnStatusModel, error) {
	var board_column_statuses []BoardColumnStatusModel
	var err error
	query := r.DB.Select("board_column_status.*")

	if err := query.Find(&board_column_statuses).Error; err != nil {
		slog.Error("failed to fetch board column statuses", "error", err)
	}
	return &board_column_statuses, err
}

func (r BoardColumnStatusRepository) FindByID(board_column_status_id uuid.UUID) (*BoardColumnStatusModel, error) {
	var board_column_status BoardColumnStatusModel
	query := r.DB.First(&board_column_status, board_column_status_id)

	if query.Error != nil {
		slog.Error("failed to find board column status by id", "error", query.Error)
	}
	return &board_column_status, query.Error
}

func (r BoardColumnStatusRepository) Create(board_column_status *BoardColumnStatusModel) error {
	query := r.DB.Create(&board_column_status)
	if query.Error != nil {
		slog.Error("failed to create board column status", "error", query.Error)
	}
	return query.Error
}

func (r BoardColumnStatusRepository) Update(board_column_status_id uuid.UUID, board_column_status *BoardColumnStatusModel) error {
	query := r.DB.Model(BoardColumnStatusModel{}).Where("id", board_column_status_id).Updates(&board_column_status)
	if query.Error != nil {
		slog.Error("failed to update board column status", "error", query.Error)
	}
	return query.Error
}

func (r BoardColumnStatusRepository) Destroy(board_column_status_id uuid.UUID) error {
	query := r.DB.Model(BoardColumnStatusModel{}).Delete(BoardColumnStatusModel{}, board_column_status_id)
	if query.Error != nil {
		slog.Error("failed to destroy board column status", "error", query.Error)
	}
	return query.Error
}
