package Repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"task1.go/task8/Testing/Domain"
)

type TaskRepoImplement struct {
	collection *mongo.Collection
}

func NewTaskRepoImplement(coll *mongo.Collection) Domain.TaskRepository {
	return &TaskRepoImplement{
		collection: coll,
	}
}

func (tr *TaskRepoImplement) AddTask(task Domain.Task) (Domain.Task, error) {
	_, err := tr.collection.InsertOne(context.Background(), task)
	return task, err
}

func (tr *TaskRepoImplement) GetTasks(Role_ string, username string) ([]Domain.Task, error) {
	var cursor *mongo.Cursor
	var err error

	
	if Role_ == "admin" {
		cursor, err = tr.collection.Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
	} else {
	
		cursor, err = tr.collection.Find(context.Background(), bson.M{"username": username})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
	}

	var tasks []Domain.Task
	for cursor.Next(context.Background()) {
		var task Domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}


func (tr *TaskRepoImplement) GetTaskByID(id string, creter string, Role_ string) (Domain.Task, error) {
	var task Domain.Task
	newId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.Task{}, err
	}
	err = tr.collection.FindOne(context.Background(), bson.M{"_id": newId}).Decode(&task)
	if err != nil {
		return Domain.Task{}, err
	}
	if Role_ != "admin" && task.Username != creter {
		return Domain.Task{}, errors.New("you are not the creater of this task")
	}
	return task, err
}

func (tr *TaskRepoImplement) GetMyTasks(username string) ([]Domain.Task, error) {
	var task []Domain.Task
	// err := tr.collection.FindOne(context.Background(), bson.M{"creater_id": username}).Decode(&task)
	cursor, err := tr.collection.Find(context.Background(), bson.M{"creater_id": username})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var t Domain.Task
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}
		task = append(task, t)
	}

	return task, err
}

func (tr *TaskRepoImplement) DeleteTask(id string) (Domain.Task, error) {
	var task Domain.Task
	newID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.Task{}, err
	}
	err = tr.collection.FindOneAndDelete(context.Background(), bson.M{"_id": newID}).Decode(&task)
	return task, err
}

func (tr *TaskRepoImplement) UpdateTask(id string, task Domain.Task) (Domain.Task, error) {
	newID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.Task{}, err
	}
	_, err = tr.collection.UpdateOne(context.Background(), bson.M{"_id": newID}, bson.M{"$set": task})
	return task, err
}
