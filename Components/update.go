package components

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatebody note
	err := json.NewDecoder(r.Body).Decode(&updatebody)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{Key: "title", Value: updatebody.Title}}
	after := options.After
	returnOptions := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: updatebody.Title}}}}
	updateResult := connectionClient.FindOneAndUpdate(context.TODO(), filter, update, &returnOptions)
	var result primitive.M
	_ = updateResult.Decode(&result)
	json.NewEncoder(w).Encode(result)

}
