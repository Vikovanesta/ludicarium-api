package db

type ExternalGameCategory struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
