package db

type GameCategory struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
