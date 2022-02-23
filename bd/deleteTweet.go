package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, userId string) error {
	cxt, calcel := context.WithTimeout(context.Background(), time.Second*15)
	defer calcel()

	db := MongoCM.Database("twittor")
	col := db.Collection("tweet")

	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objId,
		"userid": userId,
	}

	_, err := col.DeleteOne(cxt, condition)

	return err

}
