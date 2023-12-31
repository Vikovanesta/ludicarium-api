package db

import "gorm.io/gorm"

type ReleaseDateStatus struct {
	gorm.Model
	Name        string
	Description string
}
