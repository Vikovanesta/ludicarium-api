package db

import "gorm.io/gorm"

type DateCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
