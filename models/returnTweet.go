package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnTweet struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	UserId  string             `json:"userId,omitempty" bson:"userid"`
	Message string             `json:"message,omitempty" bson:"message"`
	Date    time.Time          `json:"date,omitempty" bson:"date"`
}
