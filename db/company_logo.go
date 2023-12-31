package db

import "gorm.io/gorm"

type CompanyLogo struct {
	Image Image `gorm:"embedded"`
	gorm.Model
}
