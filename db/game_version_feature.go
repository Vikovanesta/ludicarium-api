package db

type GameVersionFeature struct {
	Model
	Title       string                     `gorm:"not null" json:"title"`
	Description string                     `json:"description"`
	CategoryID  uint                       `json:"-"`
	Category    GameVersionFeatureCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Position    int                        `json:"position"`
	Values      []GameVersionFeatureValue  `gorm:"foreignKey:GameFeatureID" json:"values"`
}
