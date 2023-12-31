package db

import "gorm.io/gorm"

type SeriesType struct {
	gorm.Model
	Name        string
	Description string
}
