package main

import (
	"fmt"

	"github.com/Iretoms/hng-task-one/api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	ServeApp()
}

func ServeApp() {
	router := gin.Default()

	publicRoute := router.Group("/api")

	route.Route(publicRoute)

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}
