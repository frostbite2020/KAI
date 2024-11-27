package database

import (
	"fmt"
	"log"
	"os"

	"MsKAI/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Database connected successfully.")
}

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.City{},
		&models.Station{},
		&models.Route{},
		&models.Session{},
		&models.User{},
		&models.BookingSeat{},
		&models.Booking{},
		&models.Carriage{},
		&models.Schedule{},
		&models.Seat{},
		&models.Train{},
	)
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	fmt.Println("Database migrated successfully.")
}

func GetDB() *gorm.DB {
	return DB
}
