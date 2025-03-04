package bd

import (
	"context"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Función que obtiene los tweets de los usuarios seguidos, con paginación.
func GetTweetsFollowers(currentUserID string, page, limit int) ([]models.ReturnTweetsFollowers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("relacion")

	skip := (page - 1) * limit

	// Pipeline de agregación:
	pipeline := mongo.Pipeline{
		// Filtrar las relaciones del usuario actual.
		{{Key: "$match", Value: bson.D{{Key: "userid", Value: currentUserID}}}},
		// Lookup para traer los tweets de los usuarios seguidos.
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "tweet"},
			{Key: "localField", Value: "relationid"},
			{Key: "foreignField", Value: "userid"},
			{Key: "as", Value: "tweet"},
		}}},
		// Desenrollar el arreglo de tweets.
		{{Key: "$unwind", Value: "$tweet"}},
		// Ordenar por fecha de tweet descendente.
		{{Key: "$sort", Value: bson.D{{Key: "tweet.date", Value: -1}}}},
		// Aplicar paginación.
		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: limit}},
		// Proyectar los campos en el formato deseado.
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: "$tweet._id"},
			{Key: "userid", Value: "$tweet.userid"},
			{Key: "userrelationid", Value: "$relationid"},
			{Key: "tweet", Value: bson.D{
				{Key: "message", Value: "$tweet.message"},
				{Key: "date", Value: "$tweet.date"},
				{Key: "_id", Value: "$tweet._id"},
			}},
		}}},
	}

	cursor, err := col.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []models.ReturnTweetsFollowers

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
