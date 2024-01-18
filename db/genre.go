package db

type Genre struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
	Slug string `gorm:"unique;not null" json:"slug"`
}
