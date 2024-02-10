package main

import (
	"log"
	"net/http"
	components "todo/Components"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createNote", components.CreateNote).Methods("POST")
	// r.HandleFunc("/getNote", GetNote).Methods("POST")
	r.HandleFunc("/update", components.UpdateProfile).Methods("PUT")
	// r.HandleFunc("/delete/{id}", DeleteNote).Methods("DELETE")
	// r.HandleFunc("/getAll", GetAllNote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
