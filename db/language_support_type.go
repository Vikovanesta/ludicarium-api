package db

import "gorm.io/gorm"

type LanguageSupportType struct {
	gorm.Model
	Name string
}
