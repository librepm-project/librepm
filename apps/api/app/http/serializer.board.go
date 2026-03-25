package http

import (
	"github.com/google/uuid"

	"apps/api/app/domain"
)

type BoardColumnRequest struct {
	ID        uuid.UUID   `json:"id"`
	Label     string      `json:"label"`
	Weight    int         `json:"weight"`
	StatusIDs []uuid.UUID `json:"statusIds"`
}

type BoardRequest struct {
	ID           uuid.UUID            `json:"id"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	BoardColumns []BoardColumnRequest `json:"boardColumns"`
}

type BoardColumnResponse struct {
	ID       uuid.UUID        `json:"id"`
	Label    string           `json:"label"`
	Weight   int              `json:"weight"`
	Statuses []StatusResponse `json:"statuses"`
}

type BoardResponse struct {
	ID           uuid.UUID             `json:"id"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	BoardColumns []BoardColumnResponse `json:"boardColumns"`
}

type BoardSerializer struct{}

func (s BoardSerializer) RequestToModel(board_request BoardRequest) domain.BoardModel {
	columns := []domain.BoardColumnModel{}
	for _, col := range board_request.BoardColumns {
		statuses := []domain.BoardColumnStatusModel{}
		for _, statusID := range col.StatusIDs {
			statuses = append(statuses, domain.BoardColumnStatusModel{
				ID:       uuid.New(),
				StatusID: statusID,
			})
		}
		columns = append(columns, domain.BoardColumnModel{
			ID:                  col.ID,
			Label:               col.Label,
			Weight:              col.Weight,
			BoardColumnStatuses: statuses,
		})
	}

	return domain.BoardModel{
		Name:         board_request.Name,
		Description:  board_request.Description,
		BoardColumns: columns,
	}
}

func (s BoardSerializer) ModelToResponse(board domain.BoardModel) BoardResponse {
	columns := []BoardColumnResponse{}
	for _, col := range board.BoardColumns {
		statuses := []StatusResponse{}
		for _, colStatus := range col.BoardColumnStatuses {
			statuses = append(statuses, StatusSerializer{}.ModelToResponse(colStatus.Status))
		}
		columns = append(columns, BoardColumnResponse{
			ID:       col.ID,
			Label:    col.Label,
			Weight:   col.Weight,
			Statuses: statuses,
		})
	}

	return BoardResponse{
		ID:           board.ID,
		Name:         board.Name,
		Description:  board.Description,
		BoardColumns: columns,
	}
}

func (s BoardSerializer) ModelToResponseMany(boards []domain.BoardModel) []BoardResponse {
	serialized := []BoardResponse{}
	for _, board := range boards {
		serialized = append(serialized, s.ModelToResponse(board))
	}
	return serialized
}
