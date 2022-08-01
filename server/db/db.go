package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URI = "mongodb+srv://sunthanhan1308:annguyenzz1@mefo-music.nppnz.mongodb.net/test"
)

func ConnectDatabase() *mongo.Collection {
	log.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("songs")
	return collection
}
