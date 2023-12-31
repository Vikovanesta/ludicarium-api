package db

import "gorm.io/gorm"

type AgeRatingTitle struct {
	gorm.Model
	Name string `gorm:"not null"`
}
