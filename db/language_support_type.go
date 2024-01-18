package db

type LanguageSupportType struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
