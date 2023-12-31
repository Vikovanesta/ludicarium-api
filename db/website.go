package db

import "gorm.io/gorm"

type Website struct {
	gorm.Model
	CategoryID uint
	Category   WebsiteCategory `gorm:"foreignKey:CategoryID"`
	GameID     uint
	Trusted    bool
	URL        string
}
