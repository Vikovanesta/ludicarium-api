package db

type Language struct {
	Model
	Name       string `gorm:"not null" json:"name"`
	Locale     string `json:"locale"`
	NativeName string `json:"nativeName"`
}
