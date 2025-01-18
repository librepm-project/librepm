package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(user *UserModel) error
	Update(user_id uuid.UUID, user *UserModel) error
	FindByID(id uuid.UUID) (*UserModel, error)
	FindByEmail(email string) (*UserModel, error)
	All() (*[]UserModel, error)
	Destroy(id uuid.UUID) error
}

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) All() (*[]UserModel, error) {
	var users []UserModel
	var err error
	query := r.DB.Select("user.*")

	if err := query.Find(&users).Error; err != nil {
		fmt.Println(err)
	}

	return &users, err
}

func (r UserRepository) FindByEmail(email string) (*UserModel, error) {
	var user UserModel
	query := r.DB.Model(UserModel{}).Where(UserModel{Email: email}).Find(&user)
	return &user, query.Error
}

func (r UserRepository) FindByID(user_id uuid.UUID) (*UserModel, error) {
	var user UserModel
	query := r.DB.Where("id", user_id).Find(&user)
	return &user, query.Error
}

func (r UserRepository) Create(user *UserModel) error {
	query := r.DB.Create(&user)
	return query.Error
}

func (r UserRepository) Update(user_id uuid.UUID, user *UserModel) error {
	query := r.DB.Model(UserModel{ID: user_id}).Updates(&user)
	return query.Error
}

func (r UserRepository) Destroy(user_id uuid.UUID) error {
	query := r.DB.Delete(r.FindByID(user_id))
	return query.Error
}
