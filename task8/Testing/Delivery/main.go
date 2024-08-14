package main

import (
	"task1.go/task8/Testing/Delivery/routers"
	"task1.go/task8/Testing/database"
)

func main() {
	database.DBconnection()
	router := routers.SetupRouter()
	router.Run(":8080")
}
