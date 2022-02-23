package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twitter_clone_backEnd/bd"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina mayor a 0", http.StatusBadRequest)
		return
	}

	//pag := int64(page)

	request, correct := bd.ReadTweet(ID, page)

	if !correct {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(request)
}
