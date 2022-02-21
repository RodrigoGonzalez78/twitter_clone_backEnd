package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

func RecTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Cuerpo del tweet invalido!"+err.Error(), 400)
	}

	register := models.GrabarTweet{
		UserId:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(register)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro"+err.Error(), 400)
	}

	if !status {
		http.Error(w, "No se logro  insertar el registro", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
