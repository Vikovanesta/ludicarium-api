package db

type Company struct {
	Model
	Name                 string           `gorm:"not null" json:"name"`
	Slug                 string           `gorm:"not null" json:"slug"`
	Description          string           `json:"description"`
	Country              int              `json:"country"` //ISO 3166-1 country code
	Status               int              `json:"status"`  // TODO: What is this?
	GameCompany          []GameCompany    `gorm:"foreignKey:CompanyID" json:"game_companies"`
	StartDate            int              `json:"start_date"`
	StartDateCategoryID  uint             `json:"-"`
	StartDateCategory    DateCategory     `gorm:"foreignKey:StartDateCategoryID" json:"start_date_category"`
	ChangeDate           int              `json:"change_date"`
	ChangeDateCategoryID uint             `json:"-"`
	ChangeDateCategory   DateCategory     `gorm:"foreignKey:ChangeDateCategoryID" json:"change_date_category"`
	ChangedCompanyID     uint             `json:"-"`
	ChangedCompany       *Company         `gorm:"foreignKey:ChangedCompanyID" json:"changed_company"`
	LogoID               uint             `json:"-"`
	Logo                 CompanyLogo      `gorm:"foreignKey:LogoID" json:"logo"`
	ParentCompanyID      uint             `json:"-"`
	ParentCompany        *Company         `gorm:"foreignKey:ParentCompanyID" json:"parent_company"`
	Websites             []CompanyWebsite `gorm:"foreignKey:CompanyID" json:"websites"`
}
