package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twitter_clone_backEnd/bd"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Deve enviar el parametro de pagina mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pageTemp)

	result, status := bd.ReadAllUsers(IDUser, pag, search, typeUser)

	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
