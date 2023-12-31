package db

import "gorm.io/gorm"

type AgeRatingContentCategory struct {
	gorm.Model
	Name string `gorm:"not null"`
}
