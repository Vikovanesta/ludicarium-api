package db

import "gorm.io/gorm"

type CharacterGender struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
