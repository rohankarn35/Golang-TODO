package components

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
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
	inserted := bson.M{}
	filter := bson.D{{Key: "_id", Value: insertResult.InsertedID}}

	connectionClient.FindOne(context.TODO(), filter).Decode(&inserted)
	json.NewEncoder(w).Encode(inserted)

}
