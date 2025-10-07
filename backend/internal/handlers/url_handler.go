package handlers

import (
	"net/http"
	"urlshortener/internal/service"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func ShortenURL(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат данных"})
		return
	}

	code, err := service.CreateShortURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "произошла ошибка при генераций ссылки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": c.Request.Host + "/" + code})
}

func Redirect(c *gin.Context) {
	code := c.Param("code")
	original, err := service.GetOriginalURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ссылка с такими параметрами не найдена!"})
		return
	}

	c.Redirect(http.StatusFound, original)
}
