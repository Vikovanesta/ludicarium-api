package db

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Locale     string
	Name       string
	NativeName string
}
