package db

type AgeRatingContentCategory struct {
	Model
	Name string `gorm:"not null" json:"name"`
}
