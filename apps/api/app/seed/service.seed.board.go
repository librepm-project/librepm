package seed

import (
	"apps/api/app/domain"
	"time"

	"github.com/google/uuid"
)

func (s SeedService) createBoard(items []BoardData) error {
	for _, item := range items {
		boardID := uuid.New()
		err := s.BoardRepository.Create(&domain.BoardModel{
			ID:           boardID,
			Name:         item.Name,
			Description:  item.Description,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			BoardColumns: []domain.BoardColumnModel{},
		})
		if err != nil {
			return err
		}

		for i, col := range item.Columns {
			columnID := uuid.New()
			err = s.BoardColumnRepository.Create(&domain.BoardColumnModel{
				ID:      columnID,
				BoardID: boardID,
				Label:   col.Label,
				Weight:  i,
			})
			if err != nil {
				return err
			}

			for _, statusName := range col.Statuses {
				status, err := s.StatusRepository.FindByName(statusName)
				if err != nil {
					return err
				}
				err = s.BoardColumnStatusRepository.Create(&domain.BoardColumnStatusModel{
					ID:            uuid.New(),
					BoardColumnID: columnID,
					StatusID:      status.ID,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
