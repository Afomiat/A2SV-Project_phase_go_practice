package routers

import (
	"github.com/gin-gonic/gin"
	"task1.go/task8/Testing/Delivery/controllers"
	"task1.go/task8/Testing/Repository"
	"task1.go/task8/Testing/Usecase"
	"task1.go/task8/Testing/database"
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