package routers

import (
	"encoding/json"
	"net/http"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	status, err := bd.ModifyRegister(t, IDUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logro modificar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
