package models

//Relaciones
type Relation struct {
	UserID     string `bson:"userid" json:"userId"`
	RelationId string `bson:"relationid" json:"relationId"`
}
