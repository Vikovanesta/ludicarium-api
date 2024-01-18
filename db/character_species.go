package db

type CharacterSpecies struct {
	Model
	Name string `gorm:"unique; not null" json:"name"`
}
