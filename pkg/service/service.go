package service

import (
	"test/pkg/model"
	"test/pkg/repository"
)

type User interface {
	FindAll() ([]model.User, error)
}

type Auth interface {
	GenerateNewPair(userId string) (string, string, error)
	Refresh(userId, token string) (string, string, error)
	ParseToken(token string) (string, error)
}

type Service struct {
	User
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
		Auth: NewAuthService(repo.Auth, repo.User),
	}
}
