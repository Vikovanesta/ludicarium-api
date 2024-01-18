package db

type AlternativeName struct {
	Model
	Name    string `gorm:"not null" json:"name"`
	Comment string `json:"comment"`
	GameID  uint   `gorm:"not null" json:"-"`
}
