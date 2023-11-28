package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/RedFC/testproject/models"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to MySQL database
	dsn := os.Getenv("DATABASE_URI")
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// assigning global DB variable with connection conn
	DB = conn

	// AutoMigrate will create the table.

	// product table migration
	DB.AutoMigrate(&models.Product{})

	// user table migration
	DB.AutoMigrate(&models.User{})
}

func InitDBTest(dns string) {

	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// assigning global DB variable with connection conn
	DB = conn

	// AutoMigrate will create the table.

	// product table migration
	DB.AutoMigrate(&models.Product{})

	// user table migration
	DB.AutoMigrate(&models.User{})
}
