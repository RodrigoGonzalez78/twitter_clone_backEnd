package bd

import (
	"context"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.GrabarTweet) (string, bool, error) {
	cxt, calcel := context.WithTimeout(context.Background(), time.Second*15)
	defer calcel()

	db := MongoCM.Database("twittor")
	col := db.Collection("tweet")

	register := bson.M{
		"userid":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(cxt, register)
	if err != nil {
		return "", false, err
	}

	objId := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
