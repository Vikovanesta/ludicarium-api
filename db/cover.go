package db

type Cover struct {
	Model
	Image Image `gorm:"embedded"`
}
