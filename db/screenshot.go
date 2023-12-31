package db

import "gorm.io/gorm"

type Screenshot struct {
	gorm.Model
	Image  Image `gorm:"embedded"`
	GameID uint
}
