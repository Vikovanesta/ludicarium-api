package db

type ReleaseDate struct {
	Model
	CategoryID uint              `json:"-"`
	Category   DateCategory      `gorm:"foreignKey:CategoryID" json:"category"`
	Date       int               `json:"date"`
	Human      string            `json:"human"`
	Month      int               `json:"month"`
	Year       int               `json:"year"`
	Region     string            `json:"region"`
	GameID     uint              `json:"-"`
	StatusID   uint              `json:"-"`
	Status     ReleaseDateStatus `gorm:"foreignKey:StatusID" json:"status"`
	PlatformID uint              `json:"-"`
	Platform   Platform          `gorm:"foreignKey:PlatformID" json:"platform"`
}
