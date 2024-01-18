package db

type PlatformLogo struct {
	Model
	Image Image `gorm:"embedded"`
}
