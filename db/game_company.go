package db

type GameCompany struct {
	Model
	CompanyID  uint    `gorm:"not null" json:"-"`
	Company    Company `gorm:"foreignKey:CompanyID" json:"company"`
	Developer  bool    `json:"developer"`
	GameID     uint    `gorm:"not null" json:"-"`
	Game       Game    `gorm:"foreignKey:GameID" json:"game"`
	Porting    bool    `json:"porting"`
	Publisher  bool    `json:"publisher"`
	Supporting bool    `json:"supporting"`
}
