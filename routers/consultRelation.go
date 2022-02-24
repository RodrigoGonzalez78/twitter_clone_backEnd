package routers

import (
	"encoding/json"
	"net/http"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/models"
)

func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation

	t.UserID = IDUser
	t.RelationId = ID

	var resp models.ResponseConsultRelation

	status, err := bd.ConsultRelation(t)

	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
