package service

import "todo-app/pkg/repository"

type Service struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

func NewService(repos *repository.Repository ) *Service {
	return &Service{}
}
