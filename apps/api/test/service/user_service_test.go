package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService_All_ReturnsUsers(t *testing.T) {
	expected := &[]domain.UserModel{
		{ID: uuid.New(), Email: "alice@example.com"},
	}
	repo := &mockrepo.MockUserRepository{
		AllFunc: func() (*[]domain.UserModel, error) { return expected, nil },
	}
	svc := domain.UserService{UserRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestUserService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockUserRepository{
		AllFunc: func() (*[]domain.UserModel, error) { return nil, repoErr },
	}
	svc := domain.UserService{UserRepository: repo}

	_, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
}

func TestUserService_Show_ReturnsUser(t *testing.T) {
	userID := uuid.New()
	expected := &domain.UserModel{ID: userID, Email: "bob@example.com"}
	repo := &mockrepo.MockUserRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.UserModel, error) {
			assert.Equal(t, userID, id)
			return expected, nil
		},
	}
	svc := domain.UserService{UserRepository: repo}

	result, err := svc.Show(userID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestUserService_Update_DelegatesCorrectly(t *testing.T) {
	userID := uuid.New()
	var updatedID uuid.UUID
	repo := &mockrepo.MockUserRepository{
		UpdateFunc: func(id uuid.UUID, user *domain.UserModel) error {
			updatedID = id
			return nil
		},
	}
	svc := domain.UserService{UserRepository: repo}

	err := svc.Update(userID, &domain.UserModel{Email: "updated@example.com"})

	assert.NoError(t, err)
	assert.Equal(t, userID, updatedID)
}

func TestUserService_Destroy_DelegatesCorrectly(t *testing.T) {
	userID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockUserRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.UserService{UserRepository: repo}

	err := svc.Destroy(userID)

	assert.NoError(t, err)
	assert.Equal(t, userID, destroyedID)
}
