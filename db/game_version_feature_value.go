package db

import "gorm.io/gorm"

type GameVersionFeatureValue struct {
	gorm.Model
	GameID        uint
	Game          Game `gorm:"foreignKey:GameID"`
	GameFeatureID uint
	InclusionID   uint
	Inclusion     GameVersionFeatureInclusion `gorm:"foreignKey:InclusionID"`
	Note          string
}
