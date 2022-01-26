package bd

import (
	"context"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Insertar una funcion
func InsertRegister(u models.User) (string, bool, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("users")

	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(cxt, u)

	if err != nil {
		return "", false, err
	}

	//Obtener el id y pasar a string
	objetId, _ := result.InsertedID.(primitive.ObjectID)
	return objetId.String(), true, nil

}
