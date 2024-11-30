package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransitionRepositoryInterface interface {
	All() *[]TransitionModel
	FindByID(transition_id uuid.UUID) *TransitionModel
	Create(transition *TransitionModel)
	Update(transition_id uuid.UUID, transition *TransitionModel)
	Destroy(transition_id uuid.UUID)
}

type TransitionRepository struct {
	DB *gorm.DB
}

func (r TransitionRepository) All() *[]TransitionModel {
	var transitions []TransitionModel
	query := r.DB.Select("transition.*")

	if err := query.Find(&transitions).Error; err != nil {
		fmt.Println(err)
	}
	return &transitions
}

func (r TransitionRepository) FindByID(transition_id uuid.UUID) *TransitionModel {
	var transition TransitionModel
	fmt.Println(transition_id)
	err := r.DB.Model(TransitionModel{ID: transition_id}).Scan(&transition)

	if err != nil {
		fmt.Println(err)
	}
	return &transition
}

func (r TransitionRepository) Create(transition *TransitionModel) {
	r.DB.Create(&transition)
}

func (r TransitionRepository) Update(transition_id uuid.UUID, transition *TransitionModel) {
	r.DB.Model(TransitionModel{}).Where("id", transition_id).Updates(&transition)
}

func (r TransitionRepository) Destroy(transition_id uuid.UUID) {
	r.DB.Model(TransitionModel{}).Delete(TransitionModel{}, transition_id)
}
