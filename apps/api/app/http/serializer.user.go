package http

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
}

type SessionResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserSerializer struct{}

func (s UserSerializer) RequestToModel(user_request UserRequest) domain.UserModel {
	return domain.UserModel{
		Email:     user_request.Email,
		FirstName: user_request.FirstName,
		LastName:  user_request.LastName,
		Phone:     user_request.Phone,
		Language:  user_request.Language,
		Country:   user_request.Country,
	}
}

func (s UserSerializer) ModelToResponse(user domain.UserModel) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Language:  user.Language,
		Country:   user.Country,
	}
}

func (s UserSerializer) ModelToResponseMany(users []domain.UserModel) []UserResponse {
	var serialized []UserResponse
	for _, user := range users {
		serialized = append(serialized, s.ModelToResponse(user))
	}
	return serialized
}

func (c UserSerializer) ToSessionResponse(session domain.UserSessionCreateReturn) SessionResponse {
	return SessionResponse{
		User:  c.ModelToResponse(session.User),
		Token: session.Token,
	}
}
