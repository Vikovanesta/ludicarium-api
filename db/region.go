package db

type Region struct {
	Model
	Name       string `gorm:"not null" json:"name"`
	Category   string `json:"category"` // local, continent
	Identifier string `gorm:"unique, not null" json:"identifier"`
}
