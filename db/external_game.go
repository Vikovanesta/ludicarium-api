package db

type ExternalGame struct {
	Model
	Name       string               `json:"name"`
	UID        string               `json:"uid"`
	CategoryID uint                 `json:"-"`
	Category   ExternalGameCategory `gorm:"foreignKey:CategoryID" json:"category"`
	CountryIDs []int                `gorm:"type:integer[]" json:"country_ids"`
	GameID     int                  `json:"-"`
	Game       Game                 `gorm:"foreignKey:GameID" json:"game"`
	MediaID    uint                 `json:"-"`
	Media      ExternalGameMedia    `gorm:"foreignKey:MediaID" json:"media"`
	PlatformID uint                 `json:"-"`
	Platform   Platform             `gorm:"foreignKey:PlatformID" json:"platform"`
	Url        string               `json:"url"`
	Year       int                  `json:"year"`
}
