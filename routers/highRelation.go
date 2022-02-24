package routers

import (
	"net/http"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

//Alta relacion
func HighRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro id es requrido", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.RelationId = ID

	status, err := bd.InsertRelation(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logro insertar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
