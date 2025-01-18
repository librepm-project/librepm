package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransitionRepositoryInterface interface {
	All() (*[]TransitionModel, error)
	FindByID(transition_id uuid.UUID) (*TransitionModel, error)
	Create(transition *TransitionModel) error
	Update(transition_id uuid.UUID, transition *TransitionModel) error
	Destroy(transition_id uuid.UUID) error
}

type TransitionRepository struct {
	DB *gorm.DB
}

func (r TransitionRepository) All() (*[]TransitionModel, error) {
	var transitions []TransitionModel
	var err error
	query := r.DB.Select("transition.*")

	if err := query.Find(&transitions).Error; err != nil {
		fmt.Println(err)
	}
	return &transitions, err
}

func (r TransitionRepository) FindByID(transition_id uuid.UUID) (*TransitionModel, error) {
	var transition TransitionModel
	query := r.DB.Model(TransitionModel{ID: transition_id}).Scan(&transition)

	if query.Error != nil {
		fmt.Println(query)
	}
	return &transition, query.Error
}

func (r TransitionRepository) Create(transition *TransitionModel) error {
	query := r.DB.Create(&transition)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r TransitionRepository) Update(transition_id uuid.UUID, transition *TransitionModel) error {
	query := r.DB.Model(TransitionModel{}).Where("id", transition_id).Updates(&transition)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}

func (r TransitionRepository) Destroy(transition_id uuid.UUID) error {
	query := r.DB.Model(TransitionModel{}).Delete(TransitionModel{}, transition_id)
	if query.Error != nil {
		fmt.Println(query)
	}
	return query.Error
}
