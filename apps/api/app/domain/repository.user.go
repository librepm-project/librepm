package domain

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(user *UserModel)
	Update(user_id uuid.UUID, user *UserModel)
	FindByID(id uuid.UUID) *UserModel
	FindByEmail(email string) *UserModel
	All() *[]UserModel
	Destroy(id uuid.UUID)
}

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) All() *[]UserModel {
	var users []UserModel

	query := r.DB.Select("user.*")

	if err := query.Find(&users).Error; err != nil {
		fmt.Println(err)
	}

	return &users
}

func (r UserRepository) FindByEmail(email string) *UserModel {
	var user UserModel
	r.DB.Model(UserModel{}).Where(UserModel{Email: email}).Find(&user)
	return &user
}

func (r UserRepository) FindByID(user_id uuid.UUID) *UserModel {
	var user UserModel
	r.DB.Where("id", user_id).Find(&user)
	return &user
}

func (r UserRepository) Create(user *UserModel) {
	r.DB.Create(&user)
}

func (r UserRepository) Update(user_id uuid.UUID, user *UserModel) {
	r.DB.Model(UserModel{ID: user_id}).Updates(&user)
}

func (r UserRepository) Destroy(user_id uuid.UUID) {
	r.DB.Delete(r.FindByID(user_id))
}
