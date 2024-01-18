package db

type PlatformVersionReleaseDate struct {
	Model
	CategoryID        uint         `json:"-"`
	Category          DateCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Date              int          `json:"date"`
	Human             string       `json:"human"`
	Month             int          `json:"month"`
	Year              int          `json:"year"`
	PlatformVersionID uint         `json:"-"`
	Region            string       `json:"region"`
}
