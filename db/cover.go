package db

import "gorm.io/gorm"

type Cover struct {
	gorm.Model
	Image              Image `gorm:"embedded"`
	GameID             uint
	GameLocalizationID uint
}
