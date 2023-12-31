package db

import "gorm.io/gorm"

type LanguageSupport struct {
	gorm.Model
	GameID                uint
	LanguageID            uint
	Language              Language `gorm:"foreignKey:LanguageID"`
	LanguageSupportTypeID uint
	LanguageSupportType   LanguageSupportType `gorm:"foreignKey:LanguageSupportTypeID"`
}
