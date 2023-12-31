package db

import "gorm.io/gorm"

type GameVersionFeatureInclusion struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
