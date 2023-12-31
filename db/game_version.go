package db

import "gorm.io/gorm"

type GameVersion struct {
	gorm.Model
	Features []GameVersionFeature `gorm:"many2many:game_version_feature"`
	GameID   uint
	Game     Game   `gorm:"foreignKey:GameID"`
	Games    []Game `gorm:"many2many:game_version_game"`
}
