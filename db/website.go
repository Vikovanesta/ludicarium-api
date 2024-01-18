package db

type Website struct {
	Model
	CategoryID uint            `json:"-"`
	Category   WebsiteCategory `gorm:"foreignKey:CategoryID" json:"category"`
	GameID     uint            `json:"-"`
	Trusted    bool            `json:"trusted"`
	URL        string          `json:"url"`
}
