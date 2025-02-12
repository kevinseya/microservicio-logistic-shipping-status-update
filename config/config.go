package config

import (
	"log"
	"os"
	"strconv" // Importar strconv para la conversión

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Host     string `env:"DB_HOST"`
		Port     int    `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load() // Cargar las variables del archivo .env
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Asignar las variables de entorno
	AppConfig.Database.Host = os.Getenv("DB_HOST")
	AppConfig.Database.User = os.Getenv("DB_USER")
	AppConfig.Database.Password = os.Getenv("DB_PASSWORD")
	AppConfig.Database.Name = os.Getenv("DB_NAME")

	// Convertir DB_PORT de string a int
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int: %v", err)
	}
	AppConfig.Database.Port = port

	log.Println("Configuration loaded successfully.")
}
