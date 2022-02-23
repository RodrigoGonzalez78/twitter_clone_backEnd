package routers

import (
	"net/http"
	"twitter_clone_backEnd/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)

	if err != nil {
		http.Error(w, "Error al borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
