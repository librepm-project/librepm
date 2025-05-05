package seed

import (
	"apps/api/app/domain"
	"time"

	"github.com/google/uuid"
)

type BoardData struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Columns     []ColumnData `yaml:"columns"`
}

type ColumnData struct {
	Label string `yaml:"label"`
}

func (s SeedService) createBoard(items []BoardData) error {
	var err error
	for _, item := range items {
		err = s.BoardRepository.Create(&domain.BoardModel{
			ID:           uuid.New(),
			Name:         item.Name,
			Description:  item.Description,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			BoardColumns: []domain.BoardColumnModel{},
		})
	}
	return err
}
