package service_test

import (
	"errors"
	"os"
	"testing"

	"apps/api/app/domain"
	mockrepo "apps/api/test/mock/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAttachmentService_AllByEntity_ReturnsAttachments(t *testing.T) {
	entityID := uuid.New()
	expected := &[]domain.AttachmentModel{
		{ID: uuid.New(), Filename: "file.pdf"},
	}
	repo := &mockrepo.MockAttachmentRepository{
		AllByEntityFunc: func(entityType string, id uuid.UUID) (*[]domain.AttachmentModel, error) {
			assert.Equal(t, "issue", entityType)
			assert.Equal(t, entityID, id)
			return expected, nil
		},
	}
	svc := domain.AttachmentService{AttachmentRepository: repo}

	result, err := svc.AllByEntity("issue", entityID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestAttachmentService_FindByID_ReturnsAttachment(t *testing.T) {
	attachmentID := uuid.New()
	expected := &domain.AttachmentModel{ID: attachmentID, Filename: "img.png"}
	repo := &mockrepo.MockAttachmentRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.AttachmentModel, error) {
			assert.Equal(t, attachmentID, id)
			return expected, nil
		},
	}
	svc := domain.AttachmentService{AttachmentRepository: repo}

	result, err := svc.FindByID(attachmentID)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestAttachmentService_Create_WritesFileAndSavesToDB(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(origDir)

	var savedAttachment *domain.AttachmentModel
	repo := &mockrepo.MockAttachmentRepository{
		CreateFunc: func(attachment *domain.AttachmentModel) error {
			savedAttachment = attachment
			return nil
		},
	}
	svc := domain.AttachmentService{AttachmentRepository: repo}

	tmpFile, _ := os.CreateTemp("", "upload-*.txt")
	tmpFile.WriteString("file content")
	tmpFile.Seek(0, 0)
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	attachment := &domain.AttachmentModel{
		EntityType: "issue",
		EntityID:   uuid.New(),
		Filename:   "test.txt",
		MimeType:   "text/plain",
	}

	err := svc.Create(attachment, tmpFile)

	assert.NoError(t, err)
	assert.NotEmpty(t, savedAttachment.StorePath)
	assert.Contains(t, savedAttachment.StorePath, "test.txt")
}

func TestAttachmentService_Destroy_RemovesFromDBAndDisk(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile, _ := os.CreateTemp(tmpDir, "attachment-*.txt")
	tmpFile.WriteString("data")
	tmpFile.Close()
	storePath := tmpFile.Name()

	attachmentID := uuid.New()
	var destroyedID uuid.UUID

	repo := &mockrepo.MockAttachmentRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.AttachmentModel, error) {
			return &domain.AttachmentModel{ID: id, StorePath: storePath}, nil
		},
		DestroyFunc: func(id uuid.UUID) error {
			destroyedID = id
			return nil
		},
	}
	svc := domain.AttachmentService{AttachmentRepository: repo}

	err := svc.Destroy(attachmentID)

	assert.NoError(t, err)
	assert.Equal(t, attachmentID, destroyedID)
	_, statErr := os.Stat(storePath)
	assert.True(t, os.IsNotExist(statErr))
}

func TestAttachmentService_Destroy_ReturnsErrorWhenNotFound(t *testing.T) {
	repoErr := errors.New("not found")
	repo := &mockrepo.MockAttachmentRepository{
		FindByIDFunc: func(id uuid.UUID) (*domain.AttachmentModel, error) {
			return nil, repoErr
		},
	}
	svc := domain.AttachmentService{AttachmentRepository: repo}

	err := svc.Destroy(uuid.New())

	assert.ErrorIs(t, err, repoErr)
}

