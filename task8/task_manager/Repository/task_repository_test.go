package Repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task1.go/task8/task_manager/Domain"
	"task1.go/task8/task_manager/mocks"
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
	// Arrange
	mockRepo := mocks.NewTaskRepository(t)
	id := primitive.NewObjectID()
	expectedTask := Domain.Task{
		ID:       id,
		Title:    "Sample Task",
		Username: "user1",
		Status:   "pending",
	}

	// Set up the expectation
	mockRepo.On("GetTaskByID", id, "user1", "user").Return(expectedTask, nil)

	// Act
	task, err := mockRepo.GetTaskByID(id.Hex(), "user1", "user")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedTask, task)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
