package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("rohankarn")
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	errs := conn.Ping(context.TODO(), nil)
	if errs != nil {
		log.Fatal(err)
	}
	print("Connection successfull")
	return conn
}
