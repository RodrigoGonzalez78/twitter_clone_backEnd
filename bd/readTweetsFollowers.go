package bd

import (
	"context"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetsFollowers(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("relacion")

	skip := (page - 1) * 20

	condicion := make([]bson.M, 0)

	condicion = append(condicion, bson.M{"$match": bson.M{
		"userid": ID,
	}})

	condicion = append(condicion, bson.M{"$lookup": bson.M{
		"from":         "tweet",
		"locaField":    "userrelationid",
		"foreignField": "userid",
		"as":           "tweet",
	}})

	condicion = append(condicion, bson.M{"$unwind": "$tweet"})
	condicion = append(condicion, bson.M{"$sort": bson.M{
		"date": -1,
	}})

	condicion = append(condicion, bson.M{"$skip": skip})
	condicion = append(condicion, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condicion)

	var result []models.ReturnTweetsFollowers
	err := cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}
	return result, true
}
