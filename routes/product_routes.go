package routes

import (
	"fmt"

	"github.com/RedFC/testproject/handlers"
	"github.com/RedFC/testproject/utils"
	"github.com/gin-gonic/gin"
)

// product routes
func SetupProductRoutes(router *gin.Engine, apiVersion string) {
	// Create a group for /products endpoints
	productGroup := router.Group(fmt.Sprintf("/%s/%s", apiVersion, "products"))
	{
		productGroup.GET("/", handlers.GetProducts)
		productGroup.POST("/", utils.JwtMiddleware(), handlers.CreateProduct)
		productGroup.GET("/:id", handlers.GetProduct)
		productGroup.PUT("/:id", utils.JwtMiddleware(), handlers.UpdateProduct)
		productGroup.DELETE("/:id", utils.JwtMiddleware(), handlers.DeleteProduct)
	}

}
