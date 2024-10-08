package repository

import (
	"clean_architecture_Testing/domain"
	"clean_architecture_Testing/infrastracture"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
    collection *mongo.Collection
}

func NewUserRepositoryImpl(coll *mongo.Collection) domain.UserRepository {
    return &UserRepositoryImpl{
        collection: coll,
    }
}

func (ur *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
    user.Role = "user"
    isUserExist := ur.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Err()
    if isUserExist == nil {
        return domain.User{}, errors.New("user already exists")
    }
    _, err := ur.collection.InsertOne(context.Background(), user)
    return user, err
}

func (ur *UserRepositoryImpl) LoginUser(username string, password string) (domain.User, error) {
    var user domain.User
    err := ur.collection.FindOne(context.Background(), map[string]string{"username": username}).Decode(&user)
    if err != nil {
        return domain.User{}, err
    }
    if infrastracture.CheckPasswordHash(password, user.Password) {
        return user, nil
    }

    
    return domain.User{}, errors.New("invalid credentials")
}

func (ur *UserRepositoryImpl) GetAllUsers() ([]domain.User, error) {
    var users []domain.User
    cursor, err := ur.collection.Find(context.Background(), map[string]string{})
    if err != nil {
        return nil, err
    }
    if err = cursor.All(context.Background(), &users); err != nil {
        return nil, err
    }
    return users, nil
}


func (ur *UserRepositoryImpl) DeleteUserID(id string) (domain.User, error) {
    var user domain.User
    // chaining 
    newID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return domain.User{}, err
    }
    err = ur.collection.FindOneAndDelete(context.Background(), bson.M{"_id":newID}).Decode(&user)
    return user, err
}


func (ur *UserRepositoryImpl) GetUserByID(id string) (domain.User, error) {
    var user domain.User
    newID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return domain.User{}, err
    }
    err = ur.collection.FindOne(context.Background(), bson.M{"_id": newID}).Decode(&user)
    return user, err
}


func (ur *UserRepositoryImpl) GetMyProfile(username string) (domain.User, error) {
    var user domain.User
    err := ur.collection.FindOne(context.Background(), map[string]string{"username": username}).Decode(&user)
    return user, err
}

func (ur *UserRepositoryImpl) UpdateUser(id string, user domain.User) (domain.User, error) {
    var updatedUser domain.User
    newID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        
        return domain.User{}, err
    }



    
    err = ur.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": newID}, bson.M{"$set": user}).Decode(&updatedUser)
    
    
    return updatedUser, err
}


