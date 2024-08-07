package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Username string             `json:"username" bson:"username"`
	Status   string             `json:"status" bson:"status"`
}
