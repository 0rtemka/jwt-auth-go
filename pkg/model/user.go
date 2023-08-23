package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Email        string             `json:"email"`
	Name         string             `json:"name"`
	RefreshToken string             `json:"-" bson:"refresh_token"`
}
