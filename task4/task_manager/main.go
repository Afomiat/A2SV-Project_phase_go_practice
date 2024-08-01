package main

import (
	// "github.com/gin-gonic/gin"
	"task_manager/router"
	
)

func main() {
    r := router.SetupRouter()
    r.Run("localhost:8080")
}