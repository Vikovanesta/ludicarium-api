package db

import "gorm.io/gorm"

type GameMode struct {
	gorm.Model
	Name string
	Slug string
}
