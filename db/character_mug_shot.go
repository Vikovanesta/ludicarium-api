package db

import "gorm.io/gorm"

type CharacterMugShot struct {
	Image Image `gorm:"embedded"`
	gorm.Model
}
