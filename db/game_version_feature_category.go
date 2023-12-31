package db

import "gorm.io/gorm"

type GameVersionFeatureCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
