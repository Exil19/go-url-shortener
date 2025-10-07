package routes

import (
	"urlshortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/:code", handlers.Redirect)
	api := r.Group("/api")
	{
		api.POST("/shorten", handlers.ShortenURL)
	}
}
