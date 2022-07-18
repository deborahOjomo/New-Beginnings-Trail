package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"grail-api/endpoint"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/user", endpoint.CreateUser).Methods("POST")
	r.HandleFunc("/api/user", endpoint.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{reference}", endpoint.GetUser).Methods("GET")
	r.HandleFunc("/api/user/{reference}", endpoint.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{reference}", endpoint.DeleteUser).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
