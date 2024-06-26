package types

import "time"

type UserStore interface {
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uint) (*User, error)
	CreateUser(User) error
	UpdateUser(User) error
	DeleteUser(User) error
}

type PostStore interface {
	GetPosts() ([]Post, error)
	GetPostsByUserID(userID uint) ([]Post, error)
	CreatePost(Post) error
	UpdatePost(Post) error
	DeletePost(Post) error
}

type User struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Password string    `json:"-"`
	CratedAt time.Time `json:"created_at"`
}

type Post struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	Body             string    `json:"body"`
	Image            string    `json:"image"`
	UserID           uint      `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=30"`
}

type UpdateUserPayload struct {
	ID       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=30"`
}

type PostCreateAndUpdatePayload struct {
	Title            string `json:"title" validate:"required,min=2,max=100"`
	ShortDescription string `json:"short_description" validate:"required,min=2,max=100"`
	Body             string `json:"body" validate:"required"`
	Image            string `json:"image" validate:"required"`
}
