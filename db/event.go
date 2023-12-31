package db

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name          string
	Slug          string
	Description   string
	StartTime     int
	EndTime       int
	TimeZone      string
	EventNetworks []EventNetwork `gorm:"foreignKey:EventID"`
	Games         []Game         `gorm:"many2many:event_game;"`
	LogoID        uint
	Logo          EventLogo `gorm:"foreignKey:LogoID"`
	LiveStreamURL string
	Videos        []GameVideo `gorm:"many2many:event_video;"`
}
