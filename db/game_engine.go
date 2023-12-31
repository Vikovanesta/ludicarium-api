package db

import "gorm.io/gorm"

type GameEngine struct {
	gorm.Model
	Name                 string
	Slug                 string
	Description          string
	ProgrammingLanguages []string `gorm:"type:varchar(255)[]"`
	LogoID               uint
	Logo                 GameEngineLogo `gorm:"foreignKey:LogoID"`
	Companies            []Company      `gorm:"many2many:game_engine_company;"`
	Platforms            []Platform     `gorm:"many2many:game_engine_platform;"`
	Game                 []*Game        `gorm:"many2many:game_game_engine;"`
}
