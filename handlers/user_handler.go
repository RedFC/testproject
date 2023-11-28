package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RedFC/testproject/database"
	"github.com/RedFC/testproject/models"
	"github.com/RedFC/testproject/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Handlers for CRUD operations

func SignUp(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input data
	if err := utils.ValidateUserInput(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate user inforation from database
	var user models.User
	err := database.DB.Where("email = ?", input.Email).First(&user).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User Already Exists"})
		return
	}

	// Hash the password before saving it to the database
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Set created and updated timestamps
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	// Create a new user
	newUser := models.User{Email: input.Email, Password: hashedPassword}

	// insert user inforamtion to database
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// response return
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User Registered SuccessFully", "data": newUser})
}

func Login(c *gin.Context) {

	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate input data
	if err := utils.ValidateUserLogin(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate user inforation from database
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Hash the password using bcrypt with a cost factor of 12
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// generating JWT
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// response return
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User Logged In SuccessFully", "data": user, "token": token})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Soft delete by setting the DeletedAt field
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Error when deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User Deleted SuccessFully", "data": user})
}
