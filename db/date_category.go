package db

type DateCategory struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
