package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"
	"libs/password_utils"

	"github.com/stretchr/testify/assert"
)

func TestUserRegisterService_Create_HashesPasswordAndCreatesUser(t *testing.T) {
	var savedUser *domain.UserModel
	repo := &mockrepo.MockUserRepository{
		CreateFunc: func(user *domain.UserModel) error {
			savedUser = user
			return nil
		},
	}
	svc := domain.UserRegisterService{UserRepository: repo, UserRoleRepository: &mockrepo.MockUserRoleRepository{}}

	result, err := svc.Create("new@example.com", "plainpassword", "Jane", "Doe")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new@example.com", result.Email)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.NotEqual(t, "plainpassword", result.PasswordHash)
	assert.True(t, password_utils.CheckPasswordHash("plainpassword", result.PasswordHash))
	assert.Equal(t, savedUser, result)
}

func TestUserRegisterService_Create_ReturnsErrorOnRepoFailure(t *testing.T) {
	repoErr := errors.New("duplicate email")
	repo := &mockrepo.MockUserRepository{
		CreateFunc: func(user *domain.UserModel) error {
			return repoErr
		},
	}
	svc := domain.UserRegisterService{UserRepository: repo}

	result, err := svc.Create("dup@example.com", "pass", "A", "B")

	assert.ErrorIs(t, err, repoErr)
	assert.NotNil(t, result)
}
