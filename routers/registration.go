package routers

import (
	"encoding/json"
	"net/http"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

func Resgistration(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email es requerido!", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contaseña debe tener almenos 6 caracteres!", 400)
		return
	}

	_, encotrado, _ := bd.CheckExistUser(t.Email)

	if encotrado {
		http.Error(w, "Ya esta registrado el email!", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)

	if err != nil {
		http.Error(w, "No se pudo registrar el usuario: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo insertar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
