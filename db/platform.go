package db

type Platform struct {
	Model
	Name             string            `gorm:"not null" json:"name"`
	Slug             string            `gorm:"not null" json:"slug"`
	Abbreviation     string            `json:"abbreviation"`
	AlternativeName  string            `json:"alternative_name"`
	Summary          string            `json:"summary"`
	CategoryID       uint              `json:"-"`
	Category         PlatformCategory  `gorm:"foreignKey:CategoryID" json:"category"`
	Generation       int               `json:"generation"`
	PlatformFamilyID uint              `json:"-"`
	PlatformFamily   PlatformFamily    `gorm:"foreignKey:PlatformFamilyID" json:"platform_family"`
	Games            []*Game           `gorm:"many2many:game_platform;" json:"games"`
	Versions         []PlatformVersion `gorm:"many2many:platform_platform_version;" json:"versions"`
	Websites         []PlatformWebsite `gorm:"foreignKey:PlatformID" json:"websites"`
}
