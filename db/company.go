package db

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name                 string
	Slug                 string
	Description          string
	Country              int
	Status               int
	GameCompany          []GameCompany `gorm:"foreignKey:CompanyID"`
	StartDate            int
	StartDateCategoryID  uint
	StartDateCategory    DateCategory `gorm:"foreignKey:StartDateCategoryID"`
	ChangeDate           int
	ChangeDateCategoryID uint
	ChangeDateCategory   DateCategory `gorm:"foreignKey:ChangeDateCategoryID"`
	ChangedCompanyID     uint
	ChangedCompany       *Company `gorm:"foreignKey:ChangedCompanyID"`
	LogoID               uint
	Logo                 CompanyLogo `gorm:"foreignKey:LogoID"`
	ParentCompanyID      uint
	ParentCompany        *Company         `gorm:"foreignKey:ParentCompanyID"`
	Websites             []CompanyWebsite `gorm:"foreignKey:CompanyID"`
}
