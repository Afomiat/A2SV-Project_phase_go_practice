package controllers

import (
	"log"
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTasksById(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Received ID: %s", id)
	task, err := data.GetTaskById(id)
	if err != nil {
		log.Printf("Error fetching task by ID: %v", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	log.Printf("Task: %v", task)
	c.IndentedJSON(http.StatusOK, task)
}

func PostTasks(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	err := data.AddTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error adding task"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var update models.Task
	if err := c.BindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	err := data.UpdateTask(id, update)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Updated"})
}

func RemoveTasks(c *gin.Context) {
	id := c.Param("id")
	err := data.RemoveTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task removed"})
}
