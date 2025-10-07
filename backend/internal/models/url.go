package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ShortCode   string `json:"shortcode" gorm:"uniqueIndex"`
	OriginalUrl string `json:"originalUrl"`
}
