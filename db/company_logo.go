package db

type CompanyLogo struct {
	Model
	Image Image `gorm:"embedded"`
}
