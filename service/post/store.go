package post

import (
	"database/sql"

	"github.com/RobTov/hmblog-golang-backend/types"
)

type PostStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *PostStore {
	return &PostStore{db: db}
}

func scanRowsIntoPosts(rows *sql.Rows) (*types.Post, error) {
	posts := new(types.Post)

	err := rows.Scan(
		&posts.ID,
		&posts.Title,
		&posts.ShortDescription,
		&posts.Body,
		&posts.Image,
		&posts.UserID,
		&posts.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return posts, err
}

func (s *PostStore) GetPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts;")
	if err != nil {
		return nil, err
	}

	posts := []types.Post{}

	for rows.Next() {
		p, err := scanRowsIntoPosts(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}

	return posts, nil
}

func (s *PostStore) GetPostsByUserID(userID uint) ([]types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts WHERE id = $1;", userID)
	if err != nil {
		return nil, err
	}

	posts := []types.Post{}
	for rows.Next() {
		p, err := scanRowsIntoPosts(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}

	return posts, nil
}

func (s *PostStore) CreatePost(post types.Post) error {
	_, err := s.db.Exec(
		"INSERT INTO posts (title, short_description, body, image, user_id) VALUES ($1, $2, $3, $4, $5);",
		post.Title,
		post.ShortDescription,
		post.Body,
		post.Image,
		post.UserID,
	)

	if err != nil {
		return err
	}

	return err
}

func (s *PostStore) UpdatePost(post types.Post) error {
	_, err := s.db.Exec(
		"UPDATE posts SET title = $1, short_description = $2, body = $3, image = $4, user_id = $5 WHERE id = $6;",
		post.Title,
		post.ShortDescription,
		post.Body,
		post.Image,
		post.UserID,
		post.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) DeletePost(post types.Post) error {
	_, err := s.db.Exec("DELETE FROM posts WHERE id = $1;", post.ID)
	if err != nil {
		return err
	}

	return nil
}
