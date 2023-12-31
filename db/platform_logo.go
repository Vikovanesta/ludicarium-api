package db

import "gorm.io/gorm"

type PlatformLogo struct {
	gorm.Model
	Image Image `gorm:"embedded"`
}
