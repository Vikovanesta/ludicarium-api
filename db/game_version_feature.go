package db

import "gorm.io/gorm"

type GameVersionFeature struct {
	gorm.Model
	Title       string
	Description string
	CategoryID  uint
	Category    GameVersionFeatureCategory `gorm:"foreignKey:CategoryID"`
	Position    int
	Values      []GameVersionFeatureValue `gorm:"foreignKey:GameFeatureID"`
}
