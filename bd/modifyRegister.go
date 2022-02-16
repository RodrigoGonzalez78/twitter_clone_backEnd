package bd

import (
	"context"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}

	register["dateBirth"] = u.DateBirth
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Bibliography) > 0 {
		register["bibliography"] = u.Bibliography
	}

	if len(u.Ubication) > 0 {
		register["ubication"] = u.Ubication
	}

	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	udtString := bson.M{
		"$set": register,
	}

	objId, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{
		"_id": bson.M{"$eq": objId},
	}

	_, err := col.UpdateOne(ctx, filtro, udtString)

	if err != nil {
		return false, err
	}

	return true, nil
}
