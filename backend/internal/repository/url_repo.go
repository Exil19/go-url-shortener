package repository

import (
	"urlshortener/internal/database"
	"urlshortener/internal/models"
)

func CreateUrl(url *models.URL) error {
	return database.DB.Create(url).Error
}

func GetUrlByCode(code string) (models.URL, error) {
	var url models.URL
	err := database.DB.Where("short_code = ?", code).First(&url).Error
	return url, err
}
