package db

import "gorm.io/gorm"

type ExternalGameCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
