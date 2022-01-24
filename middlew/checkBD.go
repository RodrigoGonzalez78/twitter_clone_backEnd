package middlew

import (
	"net/http"
	"twitter_clone_backEnd/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConection() == 0 {
			http.Error(w, "Conexion perdidad con la base de datos", 500)
		}

		next.ServeHTTP(w, r)
	}
}
