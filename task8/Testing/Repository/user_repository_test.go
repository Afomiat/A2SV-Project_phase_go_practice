package Repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task1.go/task8/Testing/Domain"
	"task1.go/task8/Testing/mocks"
)

// AddUser(user User) (User, error)
// 	LoginUser(username string, password string) (User, error)
// 	GetAllUsers() ([]User, error)
// 	GetUserByID(id string) (User, error)
// 	DeleteUserID(username string) (User, error)
// }
func TestAddUser(t *testing.T){
	mockCollaction := new(mocks.UserRepository)

	user := Domain.User{
		Username:"Afomia",
		Password:"1234" ,
	}

	mockCollaction.On("AddUser", user).Return(user, nil)

	result, err := mockCollaction.AddUser(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Password, result.Password)

	mockCollaction.AssertExpectations(t) 

}


func TestDeleteUserID(t *testing.T){
	mockCollaction := new(mocks.UserRepository)
	userID := primitive.NewObjectID().Hex()
	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil{
		t.Errorf("Error converting task ID: %v", err)
		return 
	}
	user := Domain.User{
		ID:		userObjId,
		Username: "Afomia",
		Password: "1234",
	}

	mockCollaction.On("DeleteUserID", userID).Return(user, nil)
	result, err := mockCollaction.DeleteUserID(userID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockCollaction.AssertExpectations(t)

}


func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	users := []Domain.User{
		{Username: "Afomia", Password: "1234"},
		{Username: "selam", Password: "1234"},
	}

	mockRepo.On("GetAllUsers").Return(users, nil)

	result, err := mockRepo.GetAllUsers()

	assert.NoError(t, err)
	assert.ElementsMatch(t, users, result) 
	mockRepo.AssertExpectations(t)
}


func TestGetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userID := primitive.NewObjectID().Hex()
	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		t.Errorf("Error converting user ID: %v", err)
		return
	}
	user := Domain.User{
		ID:       userObjId,
		Username: "Afomia",
		Password: "1234",
	}

	mockRepo.On("GetUserByID", userID).Return(user, nil)

	result, err := mockRepo.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}


func TestLoginUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	username := "Afomia"
	password := "1234"
	user := Domain.User{
		Username: username,
		Password: password,
	}

	mockRepo.On("LoginUser", username, password).Return(user, nil)

	result, err := mockRepo.LoginUser(username, password)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}
