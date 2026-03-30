package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIssueSerializer_RequestToModel_MapsBasicFields(t *testing.T) {
	projectID := uuid.New()
	trackerID := uuid.New()
	statusID := uuid.New()
	req := apphttp.IssueRequest{
		ProjectID:   projectID,
		TrackerID:   trackerID,
		StatusID:    statusID,
		Summary:     "Test issue",
		Description: "Some description",
	}

	model := apphttp.IssueSerializer{}.RequestToModel(req)

	assert.Equal(t, projectID, model.ProjectID)
	assert.Equal(t, trackerID, model.TrackerID)
	assert.Equal(t, statusID, model.StatusID)
	assert.Equal(t, "Test issue", model.Summary)
	assert.Equal(t, "Some description", model.Description)
}

func TestIssueSerializer_RequestToModel_MapsNullableFields(t *testing.T) {
	assignedUserID := uuid.New()
	priorityID := uuid.New()
	req := apphttp.IssueRequest{
		AssignedUserID: &assignedUserID,
		PriorityID:     &priorityID,
		Summary:        "Test",
	}

	model := apphttp.IssueSerializer{}.RequestToModel(req)

	assert.NotNil(t, model.AssignedUserID)
	assert.Equal(t, assignedUserID, *model.AssignedUserID)
	assert.NotNil(t, model.PriorityID)
	assert.Equal(t, priorityID, *model.PriorityID)
}

func TestIssueSerializer_RequestToModel_NilOptionalFields(t *testing.T) {
	req := apphttp.IssueRequest{
		Summary: "Test",
	}

	model := apphttp.IssueSerializer{}.RequestToModel(req)

	assert.Nil(t, model.AssignedUserID)
	assert.Nil(t, model.ReporterUserID)
	assert.Nil(t, model.PriorityID)
}

func TestIssueSerializer_ModelToResponse_MapsBasicFields(t *testing.T) {
	issueID := uuid.New()
	model := domain.IssueModel{
		ID:          issueID,
		Key:         "TST-1",
		Summary:     "Test issue",
		Description: "Some description",
	}

	resp := apphttp.IssueSerializer{}.ModelToResponse(model)

	assert.Equal(t, issueID, resp.ID)
	assert.Equal(t, "TST-1", resp.Key)
	assert.Equal(t, "Test issue", resp.Summary)
	assert.Equal(t, "Some description", resp.Description)
}

func TestIssueSerializer_ModelToResponse_NilAssignedUser(t *testing.T) {
	model := domain.IssueModel{
		ID:           uuid.New(),
		Summary:      "Test",
		AssignedUser: nil,
	}

	resp := apphttp.IssueSerializer{}.ModelToResponse(model)

	assert.Nil(t, resp.AssignedUser)
	assert.Nil(t, resp.AssignedUserID)
}

func TestIssueSerializer_ModelToResponse_WithAssignedUser(t *testing.T) {
	userID := uuid.New()
	model := domain.IssueModel{
		ID:             uuid.New(),
		Summary:        "Test",
		AssignedUserID: &userID,
		AssignedUser: &domain.UserModel{
			ID:    userID,
			Email: "user@example.com",
		},
	}

	resp := apphttp.IssueSerializer{}.ModelToResponse(model)

	assert.NotNil(t, resp.AssignedUser)
	assert.Equal(t, userID, resp.AssignedUser.ID)
}

func TestIssueSerializer_ModelToResponseMany_MapsAllItems(t *testing.T) {
	models := []domain.IssueModel{
		{ID: uuid.New(), Key: "TST-1", Summary: "Issue 1"},
		{ID: uuid.New(), Key: "TST-2", Summary: "Issue 2"},
	}

	resp := apphttp.IssueSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, "TST-1", resp[0].Key)
	assert.Equal(t, "TST-2", resp[1].Key)
}

func TestIssueSerializer_ModelToResponseMany_ReturnsEmptySlice(t *testing.T) {
	resp := apphttp.IssueSerializer{}.ModelToResponseMany([]domain.IssueModel{})

	assert.NotNil(t, resp)
	assert.Len(t, resp, 0)
}
