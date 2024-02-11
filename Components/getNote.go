package components

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{Key: "_id", Value: _id}}
	getbody := connectionClient.FindOne(context.TODO(), filter)

	json.NewEncoder(w).Encode(getbody)

}

func GetAllNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result []primitive.M
	curr, err := connectionClient.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for curr.Next(context.TODO()) {
		var elem primitive.M
		err := curr.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, elem)
	}
	curr.Close(context.TODO())
	json.NewEncoder(w).Encode(result)
}
