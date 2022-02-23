package handlers

import (
	"log"
	"net/http"
	"os"
	"twitter_clone_backEnd/middlew"
	"twitter_clone_backEnd/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manipulators() {
	router := mux.NewRouter()

	router.HandleFunc("/registration", middlew.CheckBD(routers.Resgistration)).Methods("POST")

	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")

	router.HandleFunc("/viewprofile", middlew.CheckBD(middlew.CheckJwt(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlew.CheckBD(middlew.CheckJwt(routers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/tweet", middlew.CheckBD(middlew.CheckJwt(routers.RecTweet))).Methods("POST")
	router.HandleFunc("/readtweets", middlew.CheckBD(middlew.CheckJwt(routers.ReadTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	hamdlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, hamdlers))
}
