package db

type CharacterGender struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
