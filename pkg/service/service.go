package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
	CreateUser(user todo.User) error
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
