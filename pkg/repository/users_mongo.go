package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var users []model.User
	coll := r.db.Collection(usersCollection)
	filter := bson.D{{}}

	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return users, err
	}

	for cur.Next(context.TODO()) {
		var user model.User
		if err = cur.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UsersMongo) FindUserById(userId primitive.ObjectID) (model.User, error) {
	var user model.User
	coll := r.db.Collection(usersCollection)
	filter := bson.D{{"_id", userId}}

	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
