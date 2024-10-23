package main

import (
	"example/api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	// as default server run in localhost:8080
	router.Run()
}
