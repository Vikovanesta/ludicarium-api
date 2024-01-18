package db

type PlatformWebsite struct {
	Model
	CategoryID uint            `json:"-"`
	Category   WebsiteCategory `gorm:"foreignKey:CategoryID" json:"category"`
	PlatformID uint            `json:"-"`
	Platform   Platform        `gorm:"foreignKey:PlatformID" json:"platform"`
	Trusted    bool            `json:"trusted"`
	URL        string          `json:"url"`
}
