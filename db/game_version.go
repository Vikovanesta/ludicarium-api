package db

type GameVersion struct {
	Model
	Features []GameVersionFeature `gorm:"many2many:game_version_feature" json:"features"`
	GameID   uint                 `gorm:"not null" json:"-"`
	Game     Game                 `gorm:"foreignKey:GameID" json:"game"`
	Games    []Game               `gorm:"many2many:game_version_game" json:"games"`
}
