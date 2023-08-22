package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"test/pkg/repository"
	"time"
)

const (
	secretKey = "asKAHSDfhjklg514HKkJulkj(/KhkslKJG" // use env variables to store it
)

type AuthService struct {
	authRepo repository.Auth
	userRepo repository.User
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func NewAuthService(authRepo repository.Auth, userRepo repository.User) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *AuthService) GenerateNewPair(userId string) (string, string, error) {
	if err := s.checkUser(userId); err != nil {
		return "", "", err
	}

	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", "", err
	}

	accessToken, err := generateJWTToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return "", "", err
	}
	b64Token := base64.StdEncoding.EncodeToString([]byte(refreshToken))

	if _, err := s.authRepo.Refresh(objId, b64Token); err != nil {
		return "", "", nil
	}

	return accessToken, b64Token, nil
}

func (s *AuthService) Refresh(userId, token string) (string, string, error) {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", "", err
	}

	user, err := s.userRepo.FindUserById(objId)
	if err != nil {
		return "", "", err
	}

	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", "", err
	}

	if !matches(string(decodedToken), user.RefreshToken) {
		return "", "", errors.New("tokens don't match")
	}

	newToken, err := generateRefreshToken()
	if err != nil {
		return "", "", err
	}
	b64Token := base64.StdEncoding.EncodeToString([]byte(newToken))

	userId, err = s.authRepo.Refresh(objId, hashString(newToken))
	if err != nil {
		return "", "", err
	}

	accessToken, err := generateJWTToken(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, b64Token, nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signing algo")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", errors.New("token is invalid")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	return claims.UserId, nil
}

func generateJWTToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &TokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(secretKey))
}

func generateRefreshToken() (string, error) {
	b := make([]byte, 32)

	n := rand.NewSource(time.Now().Unix())
	r := rand.New(n)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) checkUser(userId string) error {
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	_, err = s.userRepo.FindUserById(objId)
	return err
}

func hashString(p string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func matches(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
