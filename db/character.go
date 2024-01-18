package db

type Character struct {
	Model
	Name          string           `gorm:"not null" json:"name"`
	Slug          string           `gorm:"not null" json:"slug"`
	AKAS          []string         `gorm:"type:varchar[255]" json:"akas"`
	Description   string           `gorm:"type:text" json:"description"`
	GenderID      uint             `json:"-"`
	Gender        CharacterGender  `gorm:"foreignKey:GenderID" json:"gender"`
	SpeciesID     uint             `json:"-"`
	Species       CharacterSpecies `gorm:"foreignKey:SpeciesID" json:"species"`
	OriginCountry string           `json:"origin_country"`
	MugShotID     uint             `json:"-"`
	MugShot       CharacterMugShot `gorm:"foreignKey:MugShotID" json:"mug_shot"`
	Games         []*Game          `gorm:"many2many:game_character;" json:"games"`
	IsMain        bool             `gorm:"not null" json:"is_main"`
}
