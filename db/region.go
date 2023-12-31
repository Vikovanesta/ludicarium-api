package db

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Name       string
	Category   string
	Identifier string
}
