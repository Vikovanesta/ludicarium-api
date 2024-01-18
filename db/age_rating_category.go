package db

type AgeRatingCategory struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
