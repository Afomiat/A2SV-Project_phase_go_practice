package routers

import (
	"github.com/gin-gonic/gin"
	"task1.go/task8/Testing/Delivery/controllers"
	"task1.go/task8/Testing/Infrastructure"
	"task1.go/task8/Testing/Repository"
	"task1.go/task8/Testing/Usecase"
	"task1.go/task8/Testing/database"
)

func SetupTaskRoutes(router *gin.Engine) {
	taskRepo := Repository.NewTaskRepoImplement(database.TaskCollection)
	taskUsecase := Usecase.NewTaskUsecase(taskRepo)
	taskController := controllers.NewTaskController(taskUsecase)
	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(Infrastructure.AuthMiddleware())
	{

		taskRoutes.POST("/", taskController.AddTask)
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)

	}
}