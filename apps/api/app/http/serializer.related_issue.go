package http

import (
	"apps/api/app/domain"
	"github.com/google/uuid"
)

type RelatedIssueRequest struct {
	TargetIssueId string `json:"targetIssueId"`
	Type          string `json:"type"`
}

type RelatedIssueResponse struct {
	ID                   uuid.UUID     `json:"id"`
	IssueA               IssueResponse `json:"issueA"`
	IssueB               IssueResponse `json:"issueB"`
	Type                 string        `json:"type"`
	DirectionFromCurrent string        `json:"directionFromCurrent"`
	CreatedAt            interface{}   `json:"createdAt"`
}

type RelatedIssueSerializer struct{}

func (s RelatedIssueSerializer) ModelToResponse(model *domain.RelatedIssueModel, currentIssueID uuid.UUID) RelatedIssueResponse {
	direction := s.GetDirectionFromCurrent(currentIssueID, model.IssueAID, model.IssueBID, model.Type)

	return RelatedIssueResponse{
		ID:                   model.ID,
		IssueA:               IssueSerializer{}.ModelToResponse(model.IssueA),
		IssueB:               IssueSerializer{}.ModelToResponse(model.IssueB),
		Type:                 model.Type,
		DirectionFromCurrent: direction,
		CreatedAt:            model.CreatedAt,
	}
}

func (s RelatedIssueSerializer) ModelToResponseMany(models *[]domain.RelatedIssueModel, currentIssueID uuid.UUID) []RelatedIssueResponse {
	var responses []RelatedIssueResponse
	for _, model := range *models {
		responses = append(responses, s.ModelToResponse(&model, currentIssueID))
	}
	return responses
}

func (s RelatedIssueSerializer) GetDirectionFromCurrent(currentIssueID, issueAID, issueBID uuid.UUID, relationType string) string {
	if currentIssueID == issueAID {
		return relationType
	}

	reverseMap := map[string]string{
		"blocks":      "blocked_by",
		"implements":  "implemented_by",
		"duplicates":  "is_duplicated_by",
		"related":     "related",
	}

	if reversed, ok := reverseMap[relationType]; ok {
		return reversed
	}
	return relationType
}
