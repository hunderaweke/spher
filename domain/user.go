package domain

import "context"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
type UserRepository interface {
	Create(c context.Context, user *User) error
	Update(c context.Context, user *User) error
	Delete(c context.Context, userID string)
	GetByID(c context.Context, userID string) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	Fetch(c context.Context) ([]User, error)
}

type UserUsecase interface {
	Create(c context.Context, user *User) error
	Update(c context.Context, user *User) error
	Delete(c context.Context, userID string)
	GetByID(c context.Context, userID string) (User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	Fetch(c context.Context) ([]User, error)
}
