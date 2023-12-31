package db

import "gorm.io/gorm"

type GameStatus struct {
	gorm.Model
	Name string `gorm:"unique"`
}
