package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// user routes
	SetupUserLoginAndRegisterRoutes(router)
	SetupUserOperationRoutes(router)

	// task routes
	SetupTaskRoutes(router)
	return router
}
