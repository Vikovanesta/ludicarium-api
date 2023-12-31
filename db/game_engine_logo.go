package db

import "gorm.io/gorm"

type GameEngineLogo struct {
	gorm.Model
	Image Image `gorm:"embedded"`
}
