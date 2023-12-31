package db

import "gorm.io/gorm"

type EventNetwork struct {
	gorm.Model
	EventID       uint
	Event         Event `gorm:"foreignKey:EventID"`
	NetworkTypeID uint
	NetworkType   EventNetworkType `gorm:"foreignKey:NetworkTypeID"`
	URL           string
}
