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

	if err = r.DB.Select("board.*").
		Preload("BoardColumns", func(db *gorm.DB) *gorm.DB {
			return db.Order("board_column.weight ASC")
		}).
		Preload("BoardColumns.BoardColumnStatuses.Status").
		Find(&boards).Error; err != nil {
		fmt.Println(err)
	}
	return &boards, err
}

func (r BoardRepository) FindByID(board_id uuid.UUID) (*BoardModel, error) {
	var board BoardModel
	query := r.DB.
		Preload("BoardColumns", func(db *gorm.DB) *gorm.DB {
			return db.Order("board_column.weight ASC")
		}).
		Preload("BoardColumns.BoardColumnStatuses.Status").
		First(&board, board_id)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &board, query.Error
}

func (r BoardRepository) Create(board *BoardModel) error {
	query := r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&board)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r BoardRepository) Update(board_id uuid.UUID, board *BoardModel) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Delete existing associations
		var existingColumns []BoardColumnModel
		tx.Where("board_id = ?", board_id).Find(&existingColumns)
		for _, col := range existingColumns {
			tx.Where("board_column_id = ?", col.ID).Delete(&BoardColumnStatusModel{})
		}
		tx.Where("board_id = ?", board_id).Delete(&BoardColumnModel{})

		// Update board basic fields
		if err := tx.Model(&BoardModel{}).Where("id = ?", board_id).Updates(map[string]interface{}{
			"name":        board.Name,
			"description": board.Description,
		}).Error; err != nil {
			return err
		}

		// Re-add columns and statuses
		for i := range board.BoardColumns {
			board.BoardColumns[i].BoardID = board_id
			if err := tx.Create(&board.BoardColumns[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (r BoardRepository) Destroy(board_id uuid.UUID) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var existingColumns []BoardColumnModel
		tx.Where("board_id = ?", board_id).Find(&existingColumns)
		for _, col := range existingColumns {
			tx.Where("board_column_id = ?", col.ID).Delete(&BoardColumnStatusModel{})
		}
		tx.Where("board_id = ?", board_id).Delete(&BoardColumnModel{})
		tx.Delete(&BoardModel{}, board_id)
		return nil
	})
	return err
}
