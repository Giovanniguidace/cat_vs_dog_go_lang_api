package database

import (
	"api/src/settings"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoConnect create a connection with a mongo database
func MongoConnect()(*mongo.Client, error){
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(settings.DbConnectionString))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}