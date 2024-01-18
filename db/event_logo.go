package db

type EventLogo struct {
	Model
	Image Image `gorm:"embedded"`
}
