package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"test/pkg/model"
)

const (
	usersCollection = "users"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) Refresh(userId primitive.ObjectID, token string) (string, error) {
	coll := r.db.Collection(usersCollection)
	filter := bson.D{{"_id", userId}}
	update := bson.D{{"$set", bson.D{{"refresh_token", token}}}}

	var user model.User
	err := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.ID.String(), nil
}
