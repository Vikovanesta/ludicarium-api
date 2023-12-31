package db

import "gorm.io/gorm"

type Keyword struct {
	gorm.Model
	Name string
	Slug string
}
