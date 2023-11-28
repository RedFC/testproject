package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RedFC/testproject/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwt secret from env
var secret string = os.Getenv("JWT_SECRET")

// Replace this secret key with a secure secret for production use.
var secretKey = []byte(secret)

// Middleware to verify JWT token
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		// Extract claims from the token and pass them to the next handler
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract claims"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

// function for generating token
func GenerateToken(user models.User) (string, error) {

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		// Add more claims as needed
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token expiration time (1 hour in this example)
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
