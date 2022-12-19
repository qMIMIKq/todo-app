package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type AuthMysql struct {
	db *sqlx.DB
}

func (m AuthMysql) CreateUser(user todo.User) error {
	query := fmt.Sprintf("INSERT INTO `%s` (name, username, password_hash) VALUES (?, ?, ?);", usersTable)
	if err := m.db.QueryRow(query, user.Name, user.Username, user.Password).Err(); err != nil {
		return err
	}

	return nil
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}
