package service

import (
	"crypto/rand"
	"encoding/base64"
	"urlshortener/internal/models"
	"urlshortener/internal/repository"
)

func GenerateShortCode() string {
	b := make([]byte, 4)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}

func CreateShortURL(original string) (string, error) {
	code := GenerateShortCode()
	url := models.URL{ShortCode: code, OriginalUrl: original}
	err := repository.CreateUrl(&url)
	return code, err
}

func GetOriginalURL(code string) (string, error) {
	url, err := repository.GetUrlByCode(code)
	return url.OriginalUrl, err
}
