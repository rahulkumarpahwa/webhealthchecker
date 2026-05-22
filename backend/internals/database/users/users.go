package users

import (
	"database/sql"
	"fmt"

	"github.com/rahulkumarpahwa/webhealthchecker/internals/models/users"
)

type UserStorage interface {
	CreateUser(email string, passwordHash string) error
}

type Users struct {
	DB *sql.DB
}

func (u *Users) CreateUser(email string, passwordHash string) error {
	if email == "" || passwordHash == "" {
		return fmt.Errorf("Missing Email or Password Hash!")
	}

	stmt, err := u.DB.Prepare(`INSERT INTO users (email, password_hash) VALUES ($1, $2);`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(email, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) GetUserByID(id int) (*users.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("Invalid User Id!")
	}

	stmt, err := u.DB.Prepare(`SELECT (id, email, created_at) FROM users WHERE id = $1;`)

	if err != nil {
		return nil, err
	}

	var user users.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Email, & user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
