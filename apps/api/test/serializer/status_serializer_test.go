package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStatusSerializer_RequestToModel_MapsFields(t *testing.T) {
	req := apphttp.StatusRequest{Name: "Open", Color: "#00ff00"}

	model := apphttp.StatusSerializer{}.RequestToModel(req)

	assert.Equal(t, "Open", model.Name)
	assert.Equal(t, "#00ff00", model.Color)
}

func TestStatusSerializer_ModelToResponse_MapsFields(t *testing.T) {
	id := uuid.New()
	model := domain.StatusModel{ID: id, Name: "Closed", Color: "#ff0000"}

	resp := apphttp.StatusSerializer{}.ModelToResponse(model)

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "Closed", resp.Name)
	assert.Equal(t, "#ff0000", resp.Color)
}

func TestStatusSerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.StatusModel{
		{ID: uuid.New(), Name: "Open"},
		{ID: uuid.New(), Name: "Closed"},
	}

	resp := apphttp.StatusSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
	assert.Equal(t, "Open", resp[0].Name)
}

func TestStatusSerializer_ModelToResponseMany_EmptySlice(t *testing.T) {
	resp := apphttp.StatusSerializer{}.ModelToResponseMany([]domain.StatusModel{})

	assert.Empty(t, resp)
}
