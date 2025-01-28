package database

import (
	"fmt"
	"log"
	"shippment-asignment/config"
	"shippment-asignment/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// String ConnectDB
func ConnectDB() {
	cfg := config.AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	log.Printf("Connection to MySQL: %s\n", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect with database: %v", err)
	}
	log.Println("Connection sucessfully to database.")

	// Migration
	err = DB.AutoMigrate(&models.Shipment{})
	if err != nil {
		log.Fatalf("Error to migrate table with database: %v", err)
	}
	log.Println("Migration sucessfully with database.")
}
