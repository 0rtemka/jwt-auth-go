package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) Refresh(userId primitive.ObjectID, token string) (string, error) {
	return "", nil
}
