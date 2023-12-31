package db

import "gorm.io/gorm"

type GameCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
