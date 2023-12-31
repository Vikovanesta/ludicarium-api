package db

import "gorm.io/gorm"

type AgeRatingContentDescription struct {
	gorm.Model
	CategoryID  uint
	Category    AgeRatingContentCategory `gorm:"foreignKey:CategoryID"`
	Description string
}
