package main

import (
	"os"
	"urlshortener/internal/config"
	"urlshortener/internal/database"
	"urlshortener/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()
	r := gin.Default()

	routes.InitRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
