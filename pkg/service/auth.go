package service

import (
	"crypto/sha1"
	"fmt"
	todo "todo-app"
	"todo-app/pkg/repository"
)

const salt = "dasdjfs1231dasdasd"

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user todo.User) error {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
