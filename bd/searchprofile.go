package bd

import (
	"context"
	"fmt"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Buscar el perfil del usuario
func SearchProfile(ID string) (models.User, error) {

	cxt, calcel := context.WithTimeout(context.Background(), time.Second*15)

	defer calcel()
	db := MongoCM.Database("twittor")
	col := db.Collection("users")

	var profile models.User

	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objId,
	}

	err := col.FindOne(cxt, condition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		fmt.Println("Registro no encotrado" + err.Error())
		return profile, err
	}

	return profile, nil
}
