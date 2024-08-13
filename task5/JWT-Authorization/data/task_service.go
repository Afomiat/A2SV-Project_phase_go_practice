package data

import (
	"context"
	"fmt"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	cur, err := Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var task models.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting id to ObjectID: %v", err)
		return models.Task{}, err
	}
	log.Printf("ObjectID: %v", objID)

	var task models.Task
	err = Collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&task)
	if err != nil {
		log.Printf("Error finding task: %v", err)
		return task, err
	}
	log.Printf("Found task: %v", task)
	return task, nil
}

func AddTask(task models.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := Collection.InsertOne(context.TODO(), task)
	return err
}

func UpdateTask(id string, update models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}

	updateBson := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: update.Title},
			{Key: "description", Value: update.Description},
			{Key: "due_date", Value: update.DueDate},
			{Key: "status", Value: update.Status},
		}},
	}

	_, err = Collection.UpdateOne(context.TODO(), filter, updateBson)
	return err
}

func RemoveTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := Collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objID}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", id)
	}
	return nil
}
