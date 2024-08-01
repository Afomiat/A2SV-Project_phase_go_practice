package data

import (
	"fmt"
	"task_manager/models"

	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetTasks() []models.Task {
    return tasks
}

func GetTaskById(id string) (models.Task, error) {
    for _, a := range tasks {
        if a.ID == id {
            return a, nil
        }
    }
    return models.Task{}, fmt.Errorf("task not found")
}

func AddTask(task models.Task) {
    tasks = append(tasks, task)
}

func UpdateTask(id string, update models.Task) error {
    for i, a := range tasks {
        if a.ID == id {
            if update.Description != "" {
                tasks[i].Description = update.Description
            }
            if update.Title != "" {
                tasks[i].Title = update.Title
            }
            if update.Description != ""{
                tasks[i].Status = update.Status
            }
            return nil
        }
    }
    return fmt.Errorf("task not found")
}

func RemoveTask(id string) error {
    for i, a := range tasks {
        if a.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("task not found")
}
