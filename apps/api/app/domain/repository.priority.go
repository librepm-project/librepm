package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PriorityRepositoryInterface interface {
	All() (*[]PriorityModel, error)
	FindByID(priority_id uuid.UUID) (*PriorityModel, error)
	Create(priority *PriorityModel) error
	Update(priority_id uuid.UUID, priority *PriorityModel) error
	Destroy(priority_id uuid.UUID) error
}

type PriorityRepository struct {
	DB *gorm.DB
}

func (r PriorityRepository) All() (*[]PriorityModel, error) {
	var priorities []PriorityModel
	var err error
	if err = r.DB.Find(&priorities).Error; err != nil {
		slog.Error("failed to fetch priorities", "error", err)
	}
	return &priorities, err
}

func (r PriorityRepository) FindByID(priority_id uuid.UUID) (*PriorityModel, error) {
	var priority PriorityModel
	query := r.DB.First(&priority, priority_id)
	if query.Error != nil {
		slog.Error("failed to find priority by id", "error", query.Error)
	}
	return &priority, query.Error
}

func (r PriorityRepository) Create(priority *PriorityModel) error {
	query := r.DB.Create(&priority)
	if query.Error != nil {
		slog.Error("failed to create priority", "error", query.Error)
	}
	return query.Error
}

func (r PriorityRepository) Update(priority_id uuid.UUID, priority *PriorityModel) error {
	query := r.DB.Model(PriorityModel{}).Where("id", priority_id).Updates(&priority)
	if query.Error != nil {
		slog.Error("failed to update priority", "error", query.Error)
	}
	return query.Error
}

func (r PriorityRepository) Destroy(priority_id uuid.UUID) error {
	query := r.DB.Model(PriorityModel{}).Delete(PriorityModel{}, priority_id)
	if query.Error != nil {
		slog.Error("failed to destroy priority", "error", query.Error)
	}
	return query.Error
}
