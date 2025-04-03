package domain

import (
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
		err = s.BoardRepository.Create(&BoardModel{
			ID:           uuid.New(),
			Name:         item.Name,
			Description:  item.Description,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			BoardColumns: []BoardColumnModel{},
		})
	}
	return err
}
