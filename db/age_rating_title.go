package db

type AgeRatingTitle struct {
	Model
	Name string `gorm:"not null"`
}
