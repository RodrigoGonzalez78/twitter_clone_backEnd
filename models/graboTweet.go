package models

import "time"

type GrabarTweet struct {
	UserId  string    `bson:"userid" json:"userId,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
