package db

import "gorm.io/gorm"

type Artwork struct {
	gorm.Model
	Image  Image `gorm:"embedded"`
	GameID uint
}
