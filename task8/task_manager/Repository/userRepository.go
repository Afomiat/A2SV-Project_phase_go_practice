package Repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"task1.go/task7/task_manager/Domain"
)

type UserRepoImplement struct {
	collection *mongo.Collection
}

func NewUserRepoImplementation(coll *mongo.Collection) Domain.UserRepository {
	return &UserRepoImplement{
		collection: coll,
	}
}

// AddUser implements Domain.UserRepository.
func (ur *UserRepoImplement) AddUser(user Domain.User) (Domain.User, error) {
	user.Role = "user"
	_, err := ur.collection.InsertOne(context.Background(), user)
	return user, err
}

// DeleteUserID implements Domain.UserRepository.
func (ur *UserRepoImplement) DeleteUserID(id string) (Domain.User, error) {
	var user Domain.User

	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Domain.User{}, err
	}
	err = ur.collection.FindOneAndDelete(context.Background(), bson.M{"_id": ObjID}).Decode(&user)
	return user, err

}

// GetAllUsers implements Domain.UserRepository.
func (ur *UserRepoImplement) GetAllUsers() ([]Domain.User, error) {
	var users []Domain.User
	point, err := ur.collection.Find(context.Background(), map[string]string{})
	if err != nil {
		return nil, err
	}

	if err = point.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID implements Domain.UserRepository.
func (ur *UserRepoImplement) GetUserByID(id string) (Domain.User, error) {
	var user Domain.User
	ObjID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return Domain.User{}, err
	}
	err = ur.collection.FindOne(context.Background(), bson.M{"_id": ObjID}).Decode(&user)
	return user, err
}

// LoginUser implements Domain.UserRepository.
func (ur *UserRepoImplement) LoginUser(username string, password string) (Domain.User, error) {
	var user Domain.User
	err := ur.collection.FindOne(context.Background(), map[string]string{"username": username, "password": password}).Decode(&user)

	return user, err
}
