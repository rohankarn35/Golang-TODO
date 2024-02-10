package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	MongoUri := os.Getenv("MONGO_PASS")
	clientOptions := options.Client().ApplyURI(MongoUri)
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
