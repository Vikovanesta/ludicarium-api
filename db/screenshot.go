package db

type Screenshot struct {
	Model
	Image  Image `gorm:"embedded"`
	GameID uint  `json:"-"`
}
