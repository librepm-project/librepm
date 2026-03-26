package http

import (
	"time"

	"apps/api/app/domain"
	"github.com/google/uuid"
)

type IssueWorklogRequest struct {
	Seconds  int       `json:"seconds"`
	Comment  string    `json:"comment"`
	LoggedAt time.Time `json:"loggedAt"`
}

type IssueWorklogResponse struct {
	ID       uuid.UUID    `json:"id"`
	Seconds  int          `json:"seconds"`
	Comment  string       `json:"comment"`
	LoggedAt time.Time    `json:"loggedAt"`
	User     UserResponse `json:"user"`
}

type IssueWorklogSerializer struct{}

func (s IssueWorklogSerializer) RequestToModel(req IssueWorklogRequest) domain.IssueWorklogModel {
	return domain.IssueWorklogModel{
		Seconds:  req.Seconds,
		Comment:  req.Comment,
		LoggedAt: req.LoggedAt,
	}
}

func (s IssueWorklogSerializer) ModelToResponse(w domain.IssueWorklogModel) IssueWorklogResponse {
	return IssueWorklogResponse{
		ID:       w.ID,
		Seconds:  w.Seconds,
		Comment:  w.Comment,
		LoggedAt: w.LoggedAt,
		User:     UserSerializer{}.ModelToResponse(w.User),
	}
}

func (s IssueWorklogSerializer) ModelToResponseMany(worklogs []domain.IssueWorklogModel) []IssueWorklogResponse {
	serialized := []IssueWorklogResponse{}
	for _, w := range worklogs {
		serialized = append(serialized, s.ModelToResponse(w))
	}
	return serialized
}
