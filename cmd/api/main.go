package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/walber-vaz/crud-go/cmd/api/routes"
	"github.com/walber-vaz/crud-go/configs/logger"
)

func main() {
	logger.Info("Starting API")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":5000"); err != nil {
		logger.Error("Error starting server", err)
	}
}
