package store

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
type UsersStorage struct {
	db *sql.DB
}

func (u *UsersStorage) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
	`
	err := u.db.QueryRowContext(ctx, query, user.UserName, user.Password, user.Email).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
