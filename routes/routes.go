package routes

import (
	"example/api-go/controllers"

	"github.com/gin-gonic/gin"
)

// ya ini kaya router biasa mirip di javascript
func SetupRoutes(router *gin.Engine) {
	router.GET("/user", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
}
