package db

type PlatformVersionCompany struct {
	Model
	Comment      string  `json:"comment"`
	CompanyID    uint    `gorm:"not null" json:"-"`
	Company      Company `gorm:"foreignKey:CompanyID" json:"company"`
	Developer    bool    `json:"developer"`
	Manufacturer bool    `json:"manufacturer"`
}
