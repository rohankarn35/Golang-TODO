package components

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newnote note
	err := json.NewDecoder(r.Body).Decode(&newnote)
	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := connectionClient.InsertOne(context.TODO(), newnote)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(insertResult.InsertedID)

}
