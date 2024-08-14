package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"task1.go/task8/Testing/Domain"
	"task1.go/task8/Testing/Usecase"
)

type TaskController struct {
	TaskUsecase *Usecase.TaskUsecase
}

func NewTaskController(taskUsecase *Usecase.TaskUsecase) *TaskController {
	return &TaskController{TaskUsecase: taskUsecase}
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Username = c.GetString("username")
	fmt.Println(task.Username)

	if err := tc.TaskUsecase.AddTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task created successfully"})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	crater := c.GetString("username")
	Role_ := c.GetString("role")
	// if Role_ != "admin"{

	// }

	tasks, err := tc.TaskUsecase.GetTasks(Role_, crater)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	crater := c.GetString("username")
	Role_ := c.GetString("role")
	task, err := tc.TaskUsecase.GetTaskByID(id, crater, Role_)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(id, c.GetString("username"), c.GetString("role"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if task.Username != c.GetString("username") && c.GetString("role") != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not the creater of this task"})
		return
	}

	if err := tc.TaskUsecase.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully", "task": task})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Username = c.GetString("username")
	fmt.Println(task.Username)

	if err := tc.TaskUsecase.UpdateTask(id, task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}
