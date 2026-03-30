package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPrioritySerializer_RequestToModel_MapsFields(t *testing.T) {
	req := apphttp.PriorityRequest{Name: "High", Color: "#ff0000"}

	model := apphttp.PrioritySerializer{}.RequestToModel(req)

	assert.Equal(t, "High", model.Name)
	assert.Equal(t, "#ff0000", model.Color)
}

func TestPrioritySerializer_ModelToResponse_MapsFields(t *testing.T) {
	id := uuid.New()
	model := domain.PriorityModel{ID: id, Name: "Critical", Color: "#cc0000"}

	resp := apphttp.PrioritySerializer{}.ModelToResponse(model)

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "Critical", resp.Name)
	assert.Equal(t, "#cc0000", resp.Color)
}

func TestPrioritySerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.PriorityModel{
		{ID: uuid.New(), Name: "High"},
		{ID: uuid.New(), Name: "Low"},
	}

	resp := apphttp.PrioritySerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, "High", resp[0].Name)
}
