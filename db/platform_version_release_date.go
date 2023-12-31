package db

import "gorm.io/gorm"

type PlatformVersionReleaseDate struct {
	gorm.Model
	CategoryID        uint
	Category          DateCategory `gorm:"foreignKey:CategoryID"`
	Date              int
	Human             string
	Month             int
	Year              int
	PlatformVersionID uint
	Region            string
}
