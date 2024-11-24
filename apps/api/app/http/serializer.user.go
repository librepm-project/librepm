package http

import (
	"apps/api/app/domain"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
}

type Session struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UserSerializer struct{}

func (s UserSerializer) SerializeUser(user domain.UserModel) User {
	return User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Language:  user.Language,
		Country:   user.Country,
	}
}

func (s UserSerializer) SerializeUsers(users []domain.UserModel) []User {
	var serialized []User
	for _, user := range users {
		serialized = append(serialized, s.SerializeUser(user))
	}
	return serialized
}

func (c UserSerializer) SerializeSession(session domain.UserSessionCreateReturn) Session {
	return Session{
		User:  c.SerializeUser(session.User),
		Token: session.Token,
	}
}
