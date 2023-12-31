package db

import "gorm.io/gorm"

type PlatformWebsite struct {
	gorm.Model
	CategoryID uint
	Category   WebsiteCategory `gorm:"foreignKey:CategoryID"`
	PlatformID uint
	Platform   Platform `gorm:"foreignKey:PlatformID"`
	Trusted    bool
	URL        string
}
