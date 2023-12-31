package db

import "gorm.io/gorm"

type ExternalGameMedia struct {
	gorm.Model
	Name string `gorm:"unique; not null"`
}
