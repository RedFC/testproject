package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RedFC/testproject/database"
	"github.com/RedFC/testproject/models"
	"github.com/RedFC/testproject/utils"
	"github.com/gin-gonic/gin"
)

// Handlers for CRUD operations

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error when fetching products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Products Fetched SuccessFully", "data": products})
}

func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input data
	if err := utils.ValidateProductInput(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set created and updated timestamps
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error when creating product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Products Created SuccessFully", "data": input})
}

func GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product Fetched SuccessFully", "data": product})
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input data
	if err := utils.ValidateProductInput(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.UpdatedAt = time.Now()

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error when updating product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Soft delete by setting the DeletedAt field
	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error when deleting product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product Deleted SuccessFully", "data": product})
}
