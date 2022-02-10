package jwt_p

import (
	"time"
	"twitter_clone_backEnd/models"

	"github.com/dgrijalva/jwt-go"
)

func GeneringJwt(t models.User) (string, error) {
	myClave := []byte("Twittor_clone")

	payload := jwt.MapClaims{
		"email":        t.Email,
		"name":         t.Name,
		"lastName":     t.LastName,
		"dateBirth":    t.DateBirth,
		"bibliography": t.Bibliography,
		"ubication":    t.Ubication,
		"webSite":      t.WebSite,
		"_id":          t.ID.Hex(),
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStrin, err := token.SignedString(myClave)

	if err != nil {
		return tokenStrin, err
	}

	return tokenStrin, nil
}
