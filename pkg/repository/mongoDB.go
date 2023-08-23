package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(uri, dbName string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("error initializing mongoDB: %s", err.Error())
	}

	db := client.Database(dbName)

	var res bson.M
	if err := db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&res); err != nil {
		log.Fatalf("failed connection to mongoDB: %s", err.Error())
	}

	return db
}
