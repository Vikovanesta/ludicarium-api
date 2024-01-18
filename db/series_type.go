package db

type SeriesType struct {
	Model
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `json:"description"`
}
