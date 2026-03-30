package serializer_test

import (
	"testing"
	"time"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWorklogSerializer_RequestToModel_MapsFields(t *testing.T) {
	loggedAt := time.Now()
	req := apphttp.IssueWorklogRequest{
		Seconds:  3600,
		Comment:  "Fixed the bug",
		LoggedAt: loggedAt,
	}

	model := apphttp.IssueWorklogSerializer{}.RequestToModel(req)

	assert.Equal(t, 3600, model.Seconds)
	assert.Equal(t, "Fixed the bug", model.Comment)
	assert.Equal(t, loggedAt, model.LoggedAt)
}

func TestWorklogSerializer_ModelToResponse_MapsFields(t *testing.T) {
	id := uuid.New()
	userID := uuid.New()
	loggedAt := time.Now()

	model := domain.IssueWorklogModel{
		ID:       id,
		Seconds:  1800,
		Comment:  "Review done",
		LoggedAt: loggedAt,
		User:     domain.UserModel{ID: userID, Email: "dev@example.com"},
	}

	resp := apphttp.IssueWorklogSerializer{}.ModelToResponse(model)

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, 1800, resp.Seconds)
	assert.Equal(t, "Review done", resp.Comment)
	assert.Equal(t, loggedAt, resp.LoggedAt)
	assert.Equal(t, userID, resp.User.ID)
}

func TestWorklogSerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.IssueWorklogModel{
		{ID: uuid.New(), Seconds: 3600},
		{ID: uuid.New(), Seconds: 1800},
	}

	resp := apphttp.IssueWorklogSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, 3600, resp[0].Seconds)
}
