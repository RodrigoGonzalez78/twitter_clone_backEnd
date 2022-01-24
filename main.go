package main

import (
	"log"
	"twitter_clone_backEnd/bd"
	"twitter_clone_backEnd/handlers"
)

func main() {
	if bd.CheckConection() == 0 {
		log.Fatal("Sin conexcion a la BD!")

	}
	handlers.Manipulators()
}
