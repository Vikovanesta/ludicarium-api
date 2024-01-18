package db

type CompanyWebsite struct {
	Model
	CategoryID uint            `json:"-"`
	Category   WebsiteCategory `gorm:"foreignKey:CategoryID" json:"category"`
	CompanyID  uint            `json:"-"`
	Company    Company         `gorm:"foreignKey:CompanyID" json:"company"`
	IsTrusted  bool            `json:"is_trusted"`
	URL        string          `json:"url"`
}
