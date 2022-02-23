package routers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("banner")

	var extencion = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/banners/" + IDUser + "." + extencion

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir el banner"+err.Error(), http.StatusBadRequest)
		return

	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar el banner"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User

	user.Banner = IDUser + "." + extencion

	status, err := bd.ModifyRegister(user, IDUser)

	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
