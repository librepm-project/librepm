package service_test

import (
	"errors"
	"os"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"
	"libs/password_utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("JWT_SECRET", "test-secret-key")
}

func TestUserSessionService_Create_ReturnsTokenOnValidCredentials(t *testing.T) {
	userID := uuid.New()
	plainPassword := "secret123"
	hash, _ := password_utils.HashPassword(plainPassword)

	repo := &mockrepo.MockUserRepository{
		FindByEmailFunc: func(email string) (*domain.UserModel, error) {
			assert.Equal(t, "user@example.com", email)
			return &domain.UserModel{ID: userID, Email: email, PasswordHash: hash}, nil
		},
	}
	svc := domain.UserSessionService{UserRepository: repo}

	result, err := svc.Create("user@example.com", plainPassword)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, userID, result.User.ID)
}

func TestUserSessionService_Create_ReturnsErrorOnWrongPassword(t *testing.T) {
	hash, _ := password_utils.HashPassword("correct-password")

	repo := &mockrepo.MockUserRepository{
		FindByEmailFunc: func(email string) (*domain.UserModel, error) {
			return &domain.UserModel{ID: uuid.New(), Email: email, PasswordHash: hash}, nil
		},
	}
	svc := domain.UserSessionService{UserRepository: repo}

	result, err := svc.Create("user@example.com", "wrong-password")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid credentials")
	assert.Nil(t, result)
}

func TestUserSessionService_Create_ReturnsErrorWhenUserNotFound(t *testing.T) {
	repoErr := errors.New("user not found")
	repo := &mockrepo.MockUserRepository{
		FindByEmailFunc: func(email string) (*domain.UserModel, error) {
			return nil, repoErr
		},
	}
	svc := domain.UserSessionService{UserRepository: repo}

	result, err := svc.Create("notexist@example.com", "password")

	assert.ErrorIs(t, err, repoErr)
	assert.Nil(t, result)
}
