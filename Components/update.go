package components

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type newnote struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
		ID    string `json:"_id"`
	}
	var updatebody newnote
	err := json.NewDecoder(r.Body).Decode(&updatebody)
	if err != nil {
		log.Fatal(err)
	}
	objectID, err := primitive.ObjectIDFromHex(updatebody.ID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: updatebody.Title},
		{Key: "desc", Value: updatebody.Desc},
	}}}
	_, err = connectionClient.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	updatedDocument := bson.M{}
	err = connectionClient.FindOne(context.TODO(), filter).Decode(&updatedDocument)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(updatedDocument)
}
