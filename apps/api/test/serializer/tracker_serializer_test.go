package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTrackerSerializer_RequestToModel_MapsFields(t *testing.T) {
	req := apphttp.TrackerRequest{Name: "Bug", Color: "#ff4444"}

	model := apphttp.TrackerSerializer{}.RequestToModel(req)

	assert.Equal(t, "Bug", model.Name)
	assert.Equal(t, "#ff4444", model.Color)
}

func TestTrackerSerializer_ModelToResponse_MapsFields(t *testing.T) {
	id := uuid.New()
	model := domain.TrackerModel{ID: id, Name: "Feature", Color: "#4444ff"}

	resp := apphttp.TrackerSerializer{}.ModelToResponse(model)

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "Feature", resp.Name)
	assert.Equal(t, "#4444ff", resp.Color)
}

func TestTrackerSerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.TrackerModel{
		{ID: uuid.New(), Name: "Bug"},
		{ID: uuid.New(), Name: "Task"},
	}

	resp := apphttp.TrackerSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, "Bug", resp[0].Name)
}
