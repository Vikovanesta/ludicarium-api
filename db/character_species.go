package db

import "gorm.io/gorm"

type CharacterSpecies struct {
	gorm.Model
	Name string `gorm:"unique; not null"`
}
