package db

import "gorm.io/gorm"

type Platform struct {
	gorm.Model
	Name             string
	Slug             string
	Abbreviation     string
	AlternativeName  string
	Summary          string
	CategoryID       uint
	Category         PlatformCategory `gorm:"foreignKey:CategoryID"`
	Generation       int
	PlatformFamilyID uint
	PlatformFamily   PlatformFamily    `gorm:"foreignKey:PlatformFamilyID"`
	Games            []*Game           `gorm:"many2many:game_platform;"`
	Versions         []PlatformVersion `gorm:"many2many:platform_platform_version;"`
	Websites         []PlatformWebsite `gorm:"foreignKey:PlatformID"`
}
