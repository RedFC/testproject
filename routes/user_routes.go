package routes

import (
	"fmt"

	"github.com/RedFC/testproject/handlers"
	"github.com/gin-gonic/gin"
)

// user routes
func SetupUserRoutes(router *gin.Engine, apiVersion string) {
	// Create a group for /products endpoints
	userGroup := router.Group(fmt.Sprintf("/%s/%s", apiVersion, "user"))
	{
		userGroup.POST("/login", handlers.Login)
		userGroup.POST("/register", handlers.SignUp)
	}

}
