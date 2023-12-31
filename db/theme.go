package db

import "gorm.io/gorm"

type Theme struct {
	gorm.Model
	Name string
	Slug string
}
