package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twitter_clone_backEnd/bd"
)

func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	pageParam := r.URL.Query().Get("page")
	if len(pageParam) < 1 {
		http.Error(w, "Deve enviar el prametro pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pageParam)

	if err != nil || page < 1 {
		http.Error(w, "Deve enviar el prametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, err := bd.GetTweetsFollowers(IDUser, page, 20)

	if err != nil {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
