package db

type GameFranchise struct {
	Model
	GameID      uint      `gorm:"not null" json:"-"`
	Game        Game      `gorm:"foreignKey:GameID" json:"game"`
	FranchiseID uint      `gorm:"not null" json:"-"`
	Franchise   Franchise `gorm:"foreignKey:FranchiseID" json:"franchise"`
	IsMain      bool      `json:"isMain"`
}
