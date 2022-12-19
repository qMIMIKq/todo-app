package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
	CreateUser(user todo.User) error
}

type TodoList interface {
}

type TodoItem interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
	}
}
