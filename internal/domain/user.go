package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRepository interface {
	CreateUser(user User) error
}
