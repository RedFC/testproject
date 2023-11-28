package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/RedFC/testproject/database"
	"github.com/RedFC/testproject/handlers"
	"github.com/RedFC/testproject/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var token string
var user int
var productId int

func TestProductsEndpoints(t *testing.T) {
	// Initialize the test database
	database.InitDBTest("root:test12345@tcp(172.23.112.1:3306)/products?charset=utf8mb4&parseTime=True&loc=Local")

	t.Run("Get All products", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.GET("/products", handlers.GetProducts)

		// Perform a request to the test server
		req := httptest.NewRequest("GET", "/products", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		t.Log("body", w.Body.String())
		// Assertions for the response
		assert.Equal(t, 200, w.Code)
		// Add more assertions as needed
	})

	t.Run("Create Test User", func(t *testing.T) {
		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.POST("/user/register", handlers.SignUp)

		input := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{
			Name:     "Test user",
			Email:    "test@gmail.com",
			Password: "Test12345",
		}

		jsonData, _ := json.Marshal(input)

		// Perform a request to the test server
		req := httptest.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonData))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		t.Log("body", w.Body.String())
		// Assertions for the response
		if w.Code == 409 {
			assert.Equal(t, 409, w.Code)
		} else if w.Code == 201 {
			assert.Equal(t, 201, w.Code)
		}
	})

	t.Run("Login Test User", func(t *testing.T) {
		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.POST("/user/login", handlers.Login)

		input := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{
			Email:    "test@gmail.com",
			Password: "Test12345",
		}

		jsonData, _ := json.Marshal(input)

		// Perform a request to the test server
		req := httptest.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonData))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		t.Log("body", w.Body)
		var response struct {
			Status int `json:"status"`
			Data   struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password string `json:"password"`
			} `json:"data"`
			Token string `json:"token"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Assertions for the response
		assert.Equal(t, 200, w.Code)

		token = response.Token
		user = response.Data.ID
	})

	t.Run("Create Test Product", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.POST("/products", utils.JwtMiddleware(), handlers.CreateProduct)

		input := struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
		}{
			Name:        "test product",
			Description: "test description",
			Price:       2000.0,
		}

		jsonData, _ := json.Marshal(input)

		// Perform a request to the test server
		req := httptest.NewRequest("POST", "/products", bytes.NewBuffer(jsonData))

		// Set the Authorization header with the Bearer token
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		t.Log("body", w.Body.String())

		var response struct {
			Data struct {
				ID int `json:"id"`
			} `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Assertions for the response
		assert.Equal(t, 201, w.Code)
		// Add more assertions as needed
		productId = response.Data.ID
	})

	t.Run("Get Test Product", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.GET("/products/:id", handlers.GetProduct)

		id := strconv.Itoa(productId)

		t.Log("pid", id)
		// Perform a request to the test server
		req := httptest.NewRequest("GET", "/products/"+id, nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		t.Log("body", w.Body.String())

		// Assertions for the response
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Update Test Product", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.PUT("/products/:id", utils.JwtMiddleware(), handlers.UpdateProduct)

		input := struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
		}{
			Name:        "test product update",
			Description: "test description update",
			Price:       4000.0,
		}

		jsonData, _ := json.Marshal(input)

		pid := strconv.Itoa(productId)

		// Perform a request to the test server
		req := httptest.NewRequest("PUT", "/products/"+pid, bytes.NewBuffer(jsonData))
		// Set the Authorization header with the Bearer token
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		t.Log("body", w.Body.String())

		// Assertions for the response
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Get Test Product update", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.GET("/products/:id", handlers.GetProduct)

		pid := strconv.Itoa(productId)

		// Perform a request to the test server
		req := httptest.NewRequest("GET", "/products/"+pid, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		t.Log("body", w.Body.String())

		// Assertions for the response
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Delete Test Product", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.GET("/products/:id", utils.JwtMiddleware(), handlers.DeleteProduct)

		pid := strconv.Itoa(productId)

		// Perform a request to the test server
		req := httptest.NewRequest("GET", "/products/"+pid, nil)
		// Set the Authorization header with the Bearer token
		req.Header.Set("Authorization", "Bearer "+token)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		t.Log("body", w.Body.String())

		// Assertions for the response
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Get Test Product Not Found", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.GET("/products/:id", handlers.GetProduct)

		pid := strconv.Itoa(productId)

		// Perform a request to the test server
		req := httptest.NewRequest("GET", "/products/"+pid, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		t.Log("body", w.Body.String())

		// Assertions for the response
		assert.Equal(t, 404, w.Code)
	})

	t.Run("Delete Test User", func(t *testing.T) {

		// Set up a test server with the actual database implementation
		router := gin.Default()
		router.GET("/user/:id", handlers.DeleteUser)

		uid := strconv.Itoa(user)

		// Perform a request to the test server
		req := httptest.NewRequest("GET", "/user/"+uid, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		t.Log("body", w.Body.String())

		// Assertions for the response
		assert.Equal(t, 200, w.Code)
	})

}
