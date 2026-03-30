package serializer_test

import (
	"testing"

	apphttp "apps/api/app/http"
	"apps/api/app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserSerializer_RequestToModel_MapsFields(t *testing.T) {
	req := apphttp.UserRequest{
		Email:     "alice@example.com",
		FirstName: "Alice",
		LastName:  "Smith",
		Phone:     "+1234567890",
		Language:  "en",
		Country:   "US",
	}

	model := apphttp.UserSerializer{}.RequestToModel(req)

	assert.Equal(t, "alice@example.com", model.Email)
	assert.Equal(t, "Alice", model.FirstName)
	assert.Equal(t, "Smith", model.LastName)
	assert.Equal(t, "+1234567890", model.Phone)
	assert.Equal(t, "en", model.Language)
	assert.Equal(t, "US", model.Country)
}

func TestUserSerializer_ModelToResponse_MapsFields(t *testing.T) {
	id := uuid.New()
	model := domain.UserModel{
		ID:        id,
		Email:     "bob@example.com",
		FirstName: "Bob",
		LastName:  "Jones",
		Phone:     "+9876543210",
		Language:  "hu",
		Country:   "HU",
	}

	resp := apphttp.UserSerializer{}.ModelToResponse(model)

	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "bob@example.com", resp.Email)
	assert.Equal(t, "Bob", resp.FirstName)
	assert.Equal(t, "Jones", resp.LastName)
	assert.Equal(t, "+9876543210", resp.Phone)
}

func TestUserSerializer_ModelToResponseMany_ReturnsSlice(t *testing.T) {
	models := []domain.UserModel{
		{ID: uuid.New(), Email: "a@example.com"},
		{ID: uuid.New(), Email: "b@example.com"},
	}

	resp := apphttp.UserSerializer{}.ModelToResponseMany(models)

	assert.Len(t, resp, 2)
}

func TestUserSerializer_ToSessionResponse_MapsTokenAndUser(t *testing.T) {
	userID := uuid.New()
	session := domain.UserSessionCreateReturn{
		User:  domain.UserModel{ID: userID, Email: "user@example.com"},
		Token: "test-token-123",
	}

	resp := apphttp.UserSerializer{}.ToSessionResponse(session)

	assert.Equal(t, "test-token-123", resp.Token)
	assert.Equal(t, userID, resp.User.ID)
	assert.Equal(t, "user@example.com", resp.User.Email)
}
