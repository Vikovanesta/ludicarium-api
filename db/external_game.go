package db

import "gorm.io/gorm"

type ExternalGame struct {
	gorm.Model
	Name       string
	UID        string
	CategoryID uint
	Category   ExternalGameCategory `gorm:"foreignKey:CategoryID"`
	CountryIDs []int                `gorm:"type:integer[]"`
	GameID     int
	Game       Game `gorm:"foreignKey:GameID"`
	MediaID    uint
	Media      ExternalGameMedia `gorm:"foreignKey:MediaID"`
	PlatformID uint
	Platform   Platform `gorm:"foreignKey:PlatformID"`
	Url        string
	Year       int
}
