package routers

import (
	"errors"
	"strings"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"

	"github.com/dgrijalva/jwt-go"
)

var Email string
var IDUser string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myclave := []byte("Twittor_clone")

	claim := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claim, false, "", errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claim, func(t *jwt.Token) (interface{}, error) {
		return myclave, nil
	})

	if err == nil {
		_, found, _ := bd.CheckExistUser(claim.Email)

		if found {
			Email = claim.Email
			IDUser = claim.ID.Hex()
		}

		return claim, found, IDUser, nil
	}

	if tkn.Valid {
		return claim, false, "", errors.New("token invalido")
	}

	return claim, false, "", err
}
