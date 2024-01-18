package db

type GameLocalization struct {
	Model
	CoverID  uint   `json:"-"`
	Cover    Cover  `gorm:"foreignKey:CoverID" json:"cover"`
	GameID   uint   `gorm:"not null" json:"-"`
	Name     string `json:"name"`
	RegionID uint   `json:"-"`
	Region   Region `gorm:"foreignKey:RegionID" json:"region"`
}
