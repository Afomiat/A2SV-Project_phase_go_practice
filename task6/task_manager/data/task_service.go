package data

import (
	"context"

	"log"
	"time"

	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection

func init() {
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

	taskCollection = client.Database("taskdb").Collection("tasks")
}

func GetTasks(filter bson.M) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskById(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = taskCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}




func AddTask(task models.Task) error {
	_, err := taskCollection.InsertOne(context.TODO(), task)
	return err
}
func UpdateTask(id string, task models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = taskCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": task})
	return err
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = taskCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
