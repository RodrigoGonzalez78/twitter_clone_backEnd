package bd

import (
	"context"
	"log"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweet(ID string, page int) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ReturnTweet

	condicion := bson.M{
		"userid": ID,
	}

	option := options.Find()
	option.SetLimit(20)
	option.SetSort(bson.D{{Key: "date", Value: -1}})
	option.SetSkip((int64(page) - 1) * 20)

	cursor, err := col.Find(ctx, condicion, option)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.ReturnTweet
		err := cursor.Decode(&register)

		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true
}
