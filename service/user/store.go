package user

import (
	"database/sql"
	"fmt"

	"github.com/RobTov/hmblog-golang-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Avatar,
		&user.Password,
		&user.CratedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUsers() ([]types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}

	users := []types.User{}

	for rows.Next() {
		u, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, *u)
	}

	return users, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1;", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil

}

func (s *Store) GetUserByID(id uint) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec(
		"INSERT INTO users (username, email, avatar, password) VALUES ($1, $2, $3, $4);",
		user.Username,
		user.Email,
		user.Avatar,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(user types.User) error {
	_, err := s.db.Exec(
		"UPDATE users SET username = $1, email = $2, avatar = $3, password = $4 WHERE id = $5;",
		user.Username,
		user.Email,
		user.Avatar,
		user.Password,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil

}

func (s *Store) DeleteUser(user types.User) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", user.ID)
	if err != nil {
		return err
	}

	return nil
}
