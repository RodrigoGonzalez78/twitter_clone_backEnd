package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twitter_clone_backEnd/bd"
)

func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Deve enviar el prametro pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Deve enviar el prametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.ReadTweetsFollowers(IDUser, page)

	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
