package db

type GameVersionFeatureInclusion struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
