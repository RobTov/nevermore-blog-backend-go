package types

import "time"

type UserStore interface {
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByID(id uint) (User, error)
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
	UserID           string    `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
}
