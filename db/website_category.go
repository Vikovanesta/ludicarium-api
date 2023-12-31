package db

import "gorm.io/gorm"

type WebsiteCategory struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
