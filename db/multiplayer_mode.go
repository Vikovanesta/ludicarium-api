package db

import "gorm.io/gorm"

type MultiplayerMode struct {
	gorm.Model
	Campaigncoop      bool
	Dropin            bool
	GameID            uint
	Lancoop           bool
	Offlinecoop       bool
	Offlinecoopmax    int
	Offlinemax        int
	Onlinecoop        bool
	Onlinecoopmax     int
	Onlinemax         int
	PlatformID        uint
	Platform          Platform `gorm:"foreignKey:PlatformID"`
	Splitscreen       bool
	Splitscreenonline bool
}
