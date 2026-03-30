package service_test

import (
	"errors"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBoardService_All_ReturnsBoards(t *testing.T) {
	expected := &[]domain.BoardModel{
		{ID: uuid.New(), Name: "Board A"},
	}
	repo := &mockrepo.MockBoardRepository{
		AllFunc: func() (*[]domain.BoardModel, error) { return expected, nil },
	}
	svc := domain.BoardService{BoardRepository: repo}

	result, err := svc.All()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestBoardService_All_ReturnsError(t *testing.T) {
	repoErr := errors.New("db error")
	repo := &mockrepo.MockBoardRepository{
		AllFunc: func() (*[]domain.BoardModel, error) { return nil, repoErr },
	}
	svc := domain.BoardService{BoardRepository: repo}

	result, err := svc.All()

	assert.ErrorIs(t, err, repoErr)
	assert.Nil(t, result)
}

func TestBoardService_Show_ReturnsBoard(t *testing.T) {
	boardID := uuid.New()
	expected := &domain.BoardModel{ID: boardID, Name: "My Board"}
	repo := &mockrepo.MockBoardRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.BoardModel, error) {
			assert.Equal(t, boardID, id)
			return expected, nil
		},
	}
	svc := domain.BoardService{BoardRepository: repo}

	result, err := svc.Show(boardID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestBoardService_Create_DelegatesCorrectly(t *testing.T) {
	var savedBoard *domain.BoardModel
	repo := &mockrepo.MockBoardRepository{
		CreateFunc: func(board *domain.BoardModel) error {
			savedBoard = board
			return nil
		},
	}
	svc := domain.BoardService{BoardRepository: repo}

	board := &domain.BoardModel{Name: "New Board"}
	err := svc.Create(board)

	assert.NoError(t, err)
	assert.Equal(t, board, savedBoard)
}

func TestBoardService_Update_DelegatesCorrectly(t *testing.T) {
	boardID := uuid.New()
	var updatedID uuid.UUID
	var updatedBoard *domain.BoardModel
	repo := &mockrepo.MockBoardRepository{
		UpdateFunc: func(id uuid.UUID, board *domain.BoardModel) error {
			updatedID = id
			updatedBoard = board
			return nil
		},
	}
	svc := domain.BoardService{BoardRepository: repo}

	board := &domain.BoardModel{Name: "Updated Board"}
	err := svc.Update(boardID, board)

	assert.NoError(t, err)
	assert.Equal(t, boardID, updatedID)
	assert.Equal(t, board, updatedBoard)
}

func TestBoardService_Destroy_DelegatesCorrectly(t *testing.T) {
	boardID := uuid.New()
	var destroyedID uuid.UUID
	repo := &mockrepo.MockBoardRepository{
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.BoardService{BoardRepository: repo}

	err := svc.Destroy(boardID)

	assert.NoError(t, err)
	assert.Equal(t, boardID, destroyedID)
}
