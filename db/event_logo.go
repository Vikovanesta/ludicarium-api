package db

import "gorm.io/gorm"

type EventLogo struct {
	Image Image `gorm:"embedded"`
	gorm.Model
}
