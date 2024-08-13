package routers

import (
	"github.com/gin-gonic/gin"
	"task1.go/task7/task_manager/Delivery/controllers"
	"task1.go/task7/task_manager/Repository"
	"task1.go/task7/task_manager/Usecase"
	"task1.go/task7/task_manager/database"
)

func SetupUserLoginAndRegisterRoutes(router *gin.Engine) {
	userRepo := Repository.NewUserRepoImplementation(database.UserCollection)
	userUsecase := Usecase.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)
	userRoutes := router.Group("/auth")
	{
		userRoutes.POST("/login", userController.LoginUser)
		userRoutes.POST("/register", userController.RegisterUser)
	}
}