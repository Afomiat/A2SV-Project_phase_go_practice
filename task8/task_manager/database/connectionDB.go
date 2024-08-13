package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func DBconnection(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") 
	client, err := mongo.NewClient(clientOptions)
	
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	TaskCollection = client.Database("Clean_arch").Collection("tasks")
	UserCollection = client.Database("Clean_arch").Collection("users")
}