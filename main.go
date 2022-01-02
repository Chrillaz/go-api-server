package main

import (
	"log"
	"net/http"

	"go-api-server/api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/todos", handlers.GetTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
