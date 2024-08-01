package controllers

import (
    "net/http"
    "task_manager/models"
    "task_manager/data"
    "github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    albums := data.GetTasks()
    c.IndentedJSON(http.StatusOK, albums)
}

func GetTasksById(c *gin.Context) {
    id := c.Param("id")
    album, err := data.GetTaskById(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, album)
}

func PostTasks(c *gin.Context) {
    var newAlb models.Task
    if err := c.BindJSON(&newAlb); err != nil {
        return
    }
    data.AddTask(newAlb)
	
    c.IndentedJSON(http.StatusCreated, newAlb)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var update models.Task
    if err := c.BindJSON(&update); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
