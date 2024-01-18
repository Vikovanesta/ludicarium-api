package db

type GameEngineLogo struct {
	Model
	Image Image `gorm:"embedded"`
}
