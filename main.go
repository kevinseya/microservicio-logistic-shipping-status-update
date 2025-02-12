package main

import (
	"fmt"
	"log"
	"shippment-asignment/config"
	"shippment-asignment/database"
	_ "shippment-asignment/docs"
	"shippment-asignment/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Create Shippment API
func main() {
	config.LoadConfig()
	database.ConnectDB()
	router := gin.Default()

	// Redirect to Swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger-docs/index.html")
	})
	router.Static("/docs", "./docs")
	// Endpoint to documentation of Swagger
	router.GET("/swagger-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router)

	port := 6002
	log.Printf("Microservice run on port:  %d", port)

	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Error to initialize server: %v", err)
	}
}
