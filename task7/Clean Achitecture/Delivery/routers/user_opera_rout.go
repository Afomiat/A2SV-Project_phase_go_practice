package routers

import (
	"github.com/gin-gonic/gin"
	"task1.go/task7/task_manager/Delivery/controllers"
	"task1.go/task7/task_manager/Infrastructure"
	"task1.go/task7/task_manager/Repository"
	"task1.go/task7/task_manager/Usecase"
	"task1.go/task7/task_manager/database"
)

func SetupUserOperationRoutes(router *gin.Engine) {
	userRepo := Repository.NewUserRepoImplementation(database.UserCollection)
	userUsecase := Usecase.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)
	userRoutes := router.Group("/")
	userRoutes.Use(Infrastructure.AuthMiddleware())
	{
		userRoutes.GET("/users", Infrastructure.RoleMiddleware("admin"), userController.GetAllUsers)
		userRoutes.DELETE("/users/:id", userController.DeleteUserID)
		userRoutes.GET("/users/:id", userController.GetUserByID)


	}

}