package service

import (
	"test/pkg/model"
	"test/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindAll() ([]model.User, error) {
	return nil, nil
}
