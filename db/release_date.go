package db

import "gorm.io/gorm"

type ReleaseDate struct {
	gorm.Model
	CategoryID uint
	Category   DateCategory `gorm:"foreignKey:CategoryID"`
	Date       int
	Human      string
	Month      int
	Year       int
	Region     string
	GameID     uint
	StatusID   uint
	Status     ReleaseDateStatus `gorm:"foreignKey:StatusID"`
	PlatformID uint
	Platform   Platform `gorm:"foreignKey:PlatformID"`
}
