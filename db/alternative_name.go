package db

import "gorm.io/gorm"

type AlternativeName struct {
	gorm.Model
	Name    string
	Comment string
	GameID  uint
}
