package handlers

import (
	"log"
	"net/http"

	"github.com/HamelBarrer/api-go/functions"
	"github.com/gorilla/mux"
)

// Handler es la funcion que controla las rutas
func Handler() {
	router := mux.NewRouter()

	router.HandleFunc("/user", functions.UserAdd).Methods("POST")
	router.HandleFunc("/users", functions.UserGet).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
