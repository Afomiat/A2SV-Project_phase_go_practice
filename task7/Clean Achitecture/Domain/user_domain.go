package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,min=3,max=50"`
	Password string             `json:"password" bson:"password" validate:"required,min=6,max=50"`
	Role     string             `json:"role" bson:"role"`
}

type UserRepository interface {
	AddUser(user User) (User, error)
	LoginUser(username string, password string) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	DeleteUserID(username string) (User, error)
}
