package domain

import (
	"log/slog"

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
		slog.Error("failed to fetch users", "error", err)
	}

	return &users, err
}

func (r UserRepository) FindByEmail(email string) (*UserModel, error) {
	var user UserModel
	query := r.DB.Model(UserModel{}).Where(UserModel{Email: email}).Find(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	return &user, nil
}

func (r UserRepository) FindByID(user_id uuid.UUID) (*UserModel, error) {
	var user UserModel
	query := r.DB.Where("id = ?", user_id).First(&user)
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
