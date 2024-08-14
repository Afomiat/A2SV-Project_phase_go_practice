package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
    ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Title    string             `json:"title" bson:"title"`
    Username string             `json:"username" bson:"username"`
    Status   string             `json:"status" bson:"status"`
}


type TaskRepository interface {
	AddTask(task Task) (Task, error)
	GetTasks(Role_ string, username string) ([]Task, error)
	GetTaskByID(id string, creater string, Rol_ string) (Task, error)
	DeleteTask(id string) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	

}

type TaskUsecase interface{
	AddTask(task Task) (Task, error)
	GetTasks(Role_ string, username string) ([]Task, error)
	GetTaskByID(id string, creater string, Rol_ string) (Task, error)
	DeleteTask(id string) (Task, error)
	UpdateTask(id string, task Task) (Task, error)

}
