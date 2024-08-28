package domain

const UserCollection = "users"

type User struct {
	ID       string `bson:"_id" json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}
type UserRepository interface {
	Create(user *User) (User, error)
	Update(userID string, data User) (User, error)
	Delete(userID string) error
	FetchByID(userID string) (User, error)
	FetchByEmail(email string) (User, error)
	FetchByUsername(username string) (User, error)
	Fetch() ([]User, error)
}

type UserUsecase interface {
	Register(user *User) (User, error)
	Login(email, password string) (User, error)
	Update(userID string, data User) (User, error)
	Delete(userID string) error
	FetchByID(userID string) (User, error)
	FetchByEmail(email string) (User, error)
	FetchByUsername(username string) (User, error)
	Fetch() ([]User, error)
}
