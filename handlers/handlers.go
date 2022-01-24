package handlers

import (
	"log"
	"net/http"
	"os"
	"twitter_clone_backEnd/middlew"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manipulators() {
	router := mux.NewRouter()

	router.HandleFunc("/registration", middlew.CheckBD(routers.Registration)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	hamdlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, hamdlers))
}
