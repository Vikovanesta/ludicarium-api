package db

type LanguageSupport struct {
	Model
	GameID                uint                `gorm:"not null" json:"-"`
	LanguageID            uint                `gorm:"not null" json:"-"`
	Language              Language            `gorm:"foreignKey:LanguageID" json:"language"`
	LanguageSupportTypeID uint                `json:"-"`
	LanguageSupportType   LanguageSupportType `gorm:"foreignKey:LanguageSupportTypeID" json:"language_support_type"`
}
