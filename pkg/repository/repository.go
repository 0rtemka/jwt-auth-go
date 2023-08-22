package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"test/pkg/model"
)

type User interface {
	FindAllUsers() ([]model.User, error)
	FindUserById(userId primitive.ObjectID) (model.User, error)
}

type Auth interface {
	Refresh(userId primitive.ObjectID, token string) (string, error)
}

type Repository struct {
	User
	Auth
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: NewUsersMongo(db),
		Auth: NewAuthMongo(db),
	}
}
