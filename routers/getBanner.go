package routers

import (
	"io"
	"net/http"
	"os"
	"twitter_clone_backEnd/bd"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id ", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Avatar)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}
}
