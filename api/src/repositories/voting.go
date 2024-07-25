package repositories

import (
	"api/src/models"
	"api/src/settings"
	"context"
	"time"

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
