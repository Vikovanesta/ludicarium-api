package db

type GameEngine struct {
	Model
	Name                 string         `gorm:"not null" json:"name"`
	Slug                 string         `gorm:"not null" json:"slug"`
	Description          string         `gorm:"type:text" json:"description"`
	ProgrammingLanguages []string       `gorm:"type:varchar[255]" json:"programmingLanguages"`
	LogoID               uint           `json:"-"`
	Logo                 GameEngineLogo `gorm:"foreignKey:LogoID" json:"logo"`
	Companies            []Company      `gorm:"many2many:game_engine_company;" json:"companies"`
	Platforms            []Platform     `gorm:"many2many:game_engine_platform;" json:"platforms"`
	Game                 []*Game        `gorm:"many2many:game_game_engine;" json:"games"`
}
