package bd

import (
	"twitter_clone_backEnd/models"

	"golang.org/x/crypto/bcrypt"
)

func TriedLogin(email string, pass string) (models.User, bool) {
	user, found, _ := CheckExistUser(email)

	if !found {
		return user, false
	}

	passwordBytes := []byte(pass)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
