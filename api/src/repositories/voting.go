package repositories

import (
	"api/src/models"
	"api/src/settings"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertOneVote will insert a vote into database
func InsertVote(vote models.Vote, mongo_cliente *mongo.Client) (inserted_id interface{}, err error) {
	coll := mongo_cliente.Database(settings.DbName).Collection("voting")
	vote.DateCreation = time.Now()
	result, err := coll.InsertOne(context.TODO(), vote)
	if err != nil {
		return 0, nil
	}

	return result.InsertedID, nil
}

// GetVotes get all votes into database
func GetVotes(votes []models.Vote, mongo_client *mongo.Client) ([]models.Vote, error){
	coll := mongo_client.Database(settings.DbName).Collection("voting")
	filter := bson.D{}
	result, err := coll.Find(context.TODO(),filter)
	if err != nil {
		return []models.Vote{}, err
	}

	if err = result.All(context.TODO(), &votes); err != nil {
		return []models.Vote{}, err
	}

	return votes, nil

}