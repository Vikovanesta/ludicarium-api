package db

import "gorm.io/gorm"

type GameLocalization struct {
	gorm.Model
	Cover    Cover
	GameID   uint
	Name     string
	RegionID uint
	Region   Region `gorm:"foreignKey:RegionID"`
}
