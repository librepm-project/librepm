package seed

import (
	"apps/api/app/domain"
	"time"

	"github.com/google/uuid"
)

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
