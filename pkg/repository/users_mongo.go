package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"test/pkg/model"
)

type UsersMongo struct {
	db *mongo.Database
}

func NewUsersMongo(db *mongo.Database) *UsersMongo {
	return &UsersMongo{db: db}
}

func (r *UsersMongo) FindAllUsers() ([]model.User, error) {
	return nil, nil
}
