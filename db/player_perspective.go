package db

type PlayerPerspective struct {
	Model
	Name  string  `gorm:"unique;not null" json:"name"`
	Slug  string  `gorm:"unique;not null" json:"slug"`
	Games []*Game `gorm:"many2many:game_player_perspective;" json:"games"`
}
