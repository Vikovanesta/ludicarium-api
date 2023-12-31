package db

import "gorm.io/gorm"

type AgeRatingCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
