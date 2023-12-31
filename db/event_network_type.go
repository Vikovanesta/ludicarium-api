package db

import "gorm.io/gorm"

type EventNetworkType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
