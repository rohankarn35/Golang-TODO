package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createNote", CreateNote).Methods("POST")
	r.HandleFunc("/getNote", GetNote).Methods("POST")
	r.HandleFunc("/update", UpdateProfile).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteNote).Methods("DELETE")
	r.HandleFunc("/getAll", GetAllNote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
