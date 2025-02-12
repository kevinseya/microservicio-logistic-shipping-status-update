package database

import (
	"fmt"
	"log"
	"shippment-asignment/config"
	"shippment-asignment/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// String ConnectDB
func ConnectDB() {
	cfg := config.AppConfig.Database

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	log.Printf("Connecting to PostgreSQL: %s\n", dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect with database: %v", err)
	}
	log.Println("Connection successfully to database.")

	// Migration
	err = DB.AutoMigrate(&models.Shipment{})
	if err != nil {
		log.Fatalf("Error to migrate table with database: %v", err)
	}
	log.Println("Migration successfully with database.")
}
