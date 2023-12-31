package db

import "gorm.io/gorm"

type GameCompany struct {
	gorm.Model
	CompanyID  uint
	Company    Company `gorm:"foreignKey:CompanyID"`
	Developer  bool
	GameID     uint
	Game       Game `gorm:"foreignKey:GameID"`
	Porting    bool
	Publisher  bool
	Supporting bool
}
