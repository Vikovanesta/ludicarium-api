package db

type GameVersionFeatureCategory struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
