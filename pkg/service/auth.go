package service

import "test/pkg/repository"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateNewPair(userId string) (string, string, error) {
	return "", "", nil
}

func (s *AuthService) Refresh(userId, token string) (string, string, error) {
	return "", "", nil
}

func (s *AuthService) ParseToken(token string) (string, error) {
	return "", nil
}
