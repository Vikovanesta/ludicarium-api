package db

type ReleaseDateStatus struct {
	Model
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `json:"description"`
}
