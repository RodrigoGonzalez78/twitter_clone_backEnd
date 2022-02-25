package bd

import (
	"context"
	"fmt"
	"time"
	"twitter_clone_backEnd/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("twittor")
	col := db.Collection("users")

	var results []*models.User
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip((page - 1) * 20)

	query := bson.M{
		"name": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {

		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.RelationId = s.ID.Hex()

		incluir = false
		encontrado, _ = ConsultRelation(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}

		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.RelationId == ID {
			incluir = false
		}
		if incluir {
			s.Password = ""
			s.Bibliography = ""
			s.WebSite = ""
			s.Ubication = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)

	return results, true
}
