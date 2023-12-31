package db

import "gorm.io/gorm"

type PlatformVersionCompany struct {
	gorm.Model
	Comment      string
	CompanyID    uint
	Company      Company `gorm:"foreignKey:CompanyID"`
	Developer    bool
	Manufacturer bool
}
