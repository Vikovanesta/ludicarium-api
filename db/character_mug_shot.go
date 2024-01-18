package db

type CharacterMugShot struct {
	Model
	Image Image `gorm:"embedded"`
}
