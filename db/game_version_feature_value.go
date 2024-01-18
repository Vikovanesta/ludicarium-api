package db

type GameVersionFeatureValue struct {
	Model
	GameID        uint                        `json:"-"`
	Game          Game                        `gorm:"foreignKey:GameID" json:"game"`
	GameFeatureID uint                        `json:"-"`
	InclusionID   uint                        `json:"-"`
	Inclusion     GameVersionFeatureInclusion `gorm:"foreignKey:InclusionID" json:"inclusion"`
	Note          string                      `json:"note"`
}
