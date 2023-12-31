package db

import "gorm.io/gorm"

type PlatformFamily struct {
	gorm.Model
	Name string
	Slug string
}
