package bd

import (
	"context"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
)

//Consultar relaciones
func ConsultRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("relacion")

	condition := bson.M{
		"userid":     t.UserID,
		"relationid": t.RelationId,
	}

	var result models.Relation

	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		return false, err
	}

	return true, nil
}
