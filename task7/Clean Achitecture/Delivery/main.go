package main

import (
	"task1.go/task7/task_manager/Delivery/routers"
	"task1.go/task7/task_manager/database"
)

func main() {
	database.DBconnection()
	router := routers.SetupRouter()
	router.Run(":8080")
}
