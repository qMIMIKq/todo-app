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
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `username`, `password_hash`) VALUES (?, ?, ?);", usersTable)
	if err := m.db.QueryRow(query, user.Name, user.Username, user.Password).Err(); err != nil {
		return err
	}

	return nil
}

func (m AuthMysql) GetUser(username, password string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT `id` FROM %s WHERE username=? AND password_hash=?", usersTable)
	err := m.db.Get(&user, query, username, password)

	return user, err
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}
