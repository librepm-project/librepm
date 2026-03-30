package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProjectSerializer_RequestToModel_MapsFields(t *testing.T) {
	req := apphttp.ProjectRequest{
		Name:     "My Project",
		CodeName: "MP",
	}

	model := apphttp.ProjectSerializer{}.RequestToModel(req)

	assert.Equal(t, "My Project", model.Name)
	assert.Equal(t, "MP", model.CodeName)
	assert.Nil(t, model.DefaultStatusID)
	assert.Nil(t, model.DefaultTrackerID)
}

func TestProjectSerializer_RequestToModel_ParsesDefaultStatusID(t *testing.T) {
	statusID := uuid.New()
	req := apphttp.ProjectRequest{
		Name:            "Project",
		CodeName:        "PR",
		DefaultStatusID: statusID.String(),
	}

	model := apphttp.ProjectSerializer{}.RequestToModel(req)

	assert.NotNil(t, model.DefaultStatusID)
	assert.Equal(t, statusID, *model.DefaultStatusID)
}

func TestProjectSerializer_RequestToModel_ParsesDefaultTrackerID(t *testing.T) {
	trackerID := uuid.New()
	req := apphttp.ProjectRequest{
		Name:             "Project",
		CodeName:         "PR",
		DefaultTrackerID: trackerID.String(),
	}

	model := apphttp.ProjectSerializer{}.RequestToModel(req)

	assert.NotNil(t, model.DefaultTrackerID)
	assert.Equal(t, trackerID, *model.DefaultTrackerID)
}

func TestProjectSerializer_RequestToModel_IgnoresInvalidUUID(t *testing.T) {
	req := apphttp.ProjectRequest{
		Name:            "Project",
		CodeName:        "PR",
		DefaultStatusID: "not-a-uuid",
	}

	model := apphttp.ProjectSerializer{}.RequestToModel(req)

	assert.Nil(t, model.DefaultStatusID)
}

func TestProjectSerializer_ModelToResponse_MapsFields(t *testing.T) {
	projectID := uuid.New()
	statusID := uuid.New()
	model := domain.ProjectModel{
		ID:              projectID,
		Name:            "My Project",
		CodeName:        "MP",
		DefaultStatusID: &statusID,
	}

	resp := apphttp.ProjectSerializer{}.ModelToResponse(model)

	assert.Equal(t, projectID, resp.ID)
	assert.Equal(t, "My Project", resp.Name)
	assert.Equal(t, "MP", resp.CodeName)
	assert.NotNil(t, resp.DefaultStatusID)
	assert.Equal(t, statusID, *resp.DefaultStatusID)
	assert.Nil(t, resp.DefaultTrackerID)
}

func TestProjectSerializer_ModelToResponseMany_MapsAllItems(t *testing.T) {
	models := []domain.ProjectModel{
		{ID: uuid.New(), Name: "Alpha", CodeName: "AL"},
		{ID: uuid.New(), Name: "Beta", CodeName: "BT"},
		{ID: uuid.New(), Name: "Gamma", CodeName: "GM"},
	}

	resp := apphttp.ProjectSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 3)
	assert.Equal(t, "Alpha", resp[0].Name)
	assert.Equal(t, "Beta", resp[1].Name)
	assert.Equal(t, "Gamma", resp[2].Name)
}

func TestProjectSerializer_ModelToResponseMany_ReturnsEmptySlice(t *testing.T) {
	resp := apphttp.ProjectSerializer{}.ModelToResponseMany([]domain.ProjectModel{})

	assert.NotNil(t, resp)
	assert.Len(t, resp, 0)
}
