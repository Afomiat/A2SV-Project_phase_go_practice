package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/tasks", controllers.GetTasks)
		authorized.GET("/tasks/:id", controllers.GetTaskById)
		authorized.POST("/tasks", controllers.PostTask)
		authorized.PUT("/tasks/:id", controllers.UpdateTask)
		authorized.DELETE("/tasks/:id", controllers.DeleteTask)
	}
}
