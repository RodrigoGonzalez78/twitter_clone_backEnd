package routers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")

	var extencion = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/avatars/" + IDUser + "." + extencion

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir el avatar"+err.Error(), http.StatusBadRequest)
		return

	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar eñ avatar"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User

	user.Avatar = IDUser + "." + extencion

	status, err := bd.ModifyRegister(user, IDUser)

	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en la BD"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
