package db

type WebsiteCategory struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
