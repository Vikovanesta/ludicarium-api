package db

type Artwork struct {
	Model
	Image  Image `gorm:"embedded"`
	GameID uint  `gorm:"not null" json:"-"`
}
