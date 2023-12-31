package db

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name        string
	Slug        string
	AKAS        []string `gorm:"type:varchar[255]"`
	Description string
	GenderID    uint
	Gender      CharacterGender `gorm:"foreignKey:GenderID"`
	SpeciesID   uint
	Species     CharacterSpecies `gorm:"foreignKey:SpeciesID"`
	CountryName string
	MugShotID   uint
	MugShot     CharacterMugShot `gorm:"foreignKey:MugShotID"`
	Games       []*Game          `gorm:"many2many:game_character;"`
	isMain      bool
}
