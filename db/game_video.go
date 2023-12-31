package db

import "gorm.io/gorm"

type GameVideo struct {
	gorm.Model
	GameID  uint
	Name    string
	VideoID string
}
