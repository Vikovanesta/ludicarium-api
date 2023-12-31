package db

import "gorm.io/gorm"

type PlatformVersion struct {
	gorm.Model
	Name                        string
	Slug                        string
	Summary                     string
	MainManufacturerID          uint
	MainManufacturer            PlatformVersionCompany `gorm:"foreignKey:MainManufacturerID"`
	CPU                         string
	Graphics                    string
	Storage                     string
	Memory                      string
	OS                          string
	Resolutions                 string
	Connectivity                string
	Media                       string
	Output                      string
	Sound                       string
	LogoID                      uint
	Logo                        PlatformLogo                 `gorm:"foreignKey:LogoID"`
	PlatformVersionReleaseDates []PlatformVersionReleaseDate `gorm:"foreignKey:PlatformVersionID"`
	Companies                   []PlatformVersionCompany     `gorm:"many2many:platform_platform_version_company;"`
}
