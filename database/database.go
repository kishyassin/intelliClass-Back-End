package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize database connection
func InitDatabase() (*gorm.DB, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file, trying system environment variables.")
	}

	// Get database credentials from environment variables
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing database environment variables")
	}

	// Build DSN (Data Source Name) with SSL mode required
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	log.Println("Connecting to database at", host)

	// Connect to PostgreSQL using GORM
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully.")
	return db, nil
}