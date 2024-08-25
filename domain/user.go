package domain

import (
	"context"
)

type User struct {
	ID       string `bson:"_id" json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
type UserRepository interface {
	Create(user *User) (User, error)
	Update(user *User) (User, error)
	Delete(userID string)
	FetchByID(userID string) (User, error)
	FetchByEmail(email string) (User, error)
	FetchByUsername(username string) (User, error)
	Fetch(c context.Context) ([]User, error)
}

type UserUsecase interface {
	Create(user *User) (User, error)
	Update(user *User) (User, error)
	Delete(userID string)
	FetchByID(userID string) (User, error)
	FetchByEmail(email string) (User, error)
	FetchByUsername(username string) (User, error)
	Fetch(c context.Context) ([]User, error)
}
