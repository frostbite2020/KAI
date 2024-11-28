package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"MsKAI/internal/auth"
	"MsKAI/internal/database"
	"MsKAI/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "" || os.Getenv("ENV") == "development" {
		err := godotenv.Load("../../.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	// Inisialisasi auth dan database
	auth.NewAuth()
	database.InitializeDB()

	// Jalankan migrasi jika diaktifkan
	migrationStatus := os.Getenv("MIGRATION_DB")
	if migrationStatus == "TRUE" {
		log.Println("Running database migration...")
		database.MigrateDB()
	}

	// Daftarkan semua rute
	r := routes.RegisterRoutes()

	// Jalankan server
	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
