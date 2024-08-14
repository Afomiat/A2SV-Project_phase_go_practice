package Repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task1.go/task8/Testing/Domain"
	"task1.go/task8/Testing/mocks"
)

func TestAddTask(t *testing.T) {
	mockCollaction := new(mocks.TaskRepository)

	task := Domain.Task{
		Title: "Test Task",
		//   Description: "Test Description",
		Status:   "Pending",
		Username: "user1",
	}

	mockCollaction.On("AddTask", task).Return(task, nil)

	result, err := mockCollaction.AddTask(task)

	assert.NoError(t, err)
	assert.Equal(t, task.Title, result.Title)
	// assert.Equal(t, task.Description, result.Description)
	assert.Equal(t, task.Status, result.Status)
	assert.Equal(t, task.Username, result.Username)

	mockCollaction.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	// Arrange
	// mockRepo := mocks.NewTaskRepository(t)
	id := primitive.NewObjectID()
	newId := id.Hex()
	task := Domain.Task{
		ID:       id,
		Title:    "Sample Task",
		Username: "user1",
		Status:   "pending",
	}

	// // Set up the expectation
	mockRepo.On("GetTaskByID",newId, "user1", "user").Return(task, nil)

	// // Act
	result, err := mockRepo.GetTaskByID(newId, "user1", "user")
 
	// // Assert
	assert.NoError(t, err)
	assert.Equal(t, result, task)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
// GetTasks(Role_ string, username string) ([]Task, error)
func TestGetTasks(t *testing.T){
	mockRepo := new(mocks.TaskRepository)

	tasks := []Domain.Task{  
        {  
            Title:    "Test Title 1",  
            Username: "Afi",  
            Status:   "Done",  
        },  
        {  
            Title:    "Test Title 2",  
            Username: "Afi",  
            Status:   "In Progress",  
        },  
        {  
            Title:    "Test Title 3",  
            Username: "Afi",  
            Status:   "Pending",  
        }, 
		
    }  
	
	mockRepo.On("GetTasks", "user", "Afi").Return(tasks, nil)

	result, err := mockRepo.GetTasks("user", "Afi")

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)

	mockRepo.AssertExpectations(t)

}


func TestDeleteTask(t *testing.T){
	taskID := primitive.NewObjectID().Hex()
	ObjectID, err := primitive.ObjectIDFromHex(taskID)
	mockRepo := new(mocks.TaskRepository)

	task := Domain.Task{
		ID: ObjectID,
		Title:    "Sample Task",
		Username: "user1",
		Status:   "Done",
	}
	
	mockRepo.On("DeleteTask", taskID).Return(task, nil)
	

	result, err := mockRepo.DeleteTask(taskID)
	assert.NoError(t, err)
	assert.Equal(t, task, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T){
	taskID := primitive.NewObjectID().Hex()
	ObjectID, err := primitive.ObjectIDFromHex(taskID)
	mockRepo := new(mocks.TaskRepository)

	task := Domain.Task{
		ID: ObjectID,
		Title:    "Sample Task",
		Username: "user1",
		Status:   "Done",
	}
	
	mockRepo.On("UpdateTask", taskID, task).Return(task, nil)
	

	result, err := mockRepo.UpdateTask(taskID, task)
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	
	mockRepo.AssertExpectations(t)
}