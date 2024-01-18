package db

type GameStatus struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
