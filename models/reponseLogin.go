package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseLogin struct {
	UserId primitive.ObjectID `json:"user_id,omitempty"`
	Token  string             `json:"token,omitempty"`
}
