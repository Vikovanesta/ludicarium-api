package db

import "gorm.io/gorm"

type PlayerPerspective struct {
	gorm.Model
	Name  string
	Slug  string
	Games []*Game `gorm:"many2many:game_player_perspective;"`
}
