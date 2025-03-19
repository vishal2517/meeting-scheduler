package main

import (
	"meeting-scheduler/db"
	"meeting-scheduler/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "meeting-scheduler/docs"
)

// @title Meeting Scheduler API
// @version 1.0
// @description API for scheduling meetings across different time zones.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load()

	db.ConnectDatabase()
	db.ConnectRedis()

	router := gin.Default()
	routes.SetupRoutes(router)

	// Correct usage of Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
