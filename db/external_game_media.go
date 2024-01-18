package db

type ExternalGameMedia struct {
	Model
	Name string `gorm:"unique; not null" json:"name"`
}
