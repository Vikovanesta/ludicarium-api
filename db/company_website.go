package db

import "gorm.io/gorm"

type CompanyWebsite struct {
	gorm.Model
	CategoryID uint
	Category   WebsiteCategory `gorm:"foreignKey:CategoryID"`
	CompanyID  uint
	Company    Company `gorm:"foreignKey:CompanyID"`
	Trusted    bool
	URL        string
}
