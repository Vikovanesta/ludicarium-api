package db

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name                  string
	Slug                  string
	CategoryID            uint
	Category              GameCategory `gorm:"foreignKey:CategoryID"`
	FirstReleaseDate      int
	ParentGameID          uint
	ParentGame            *Game `gorm:"foreignKey:ParentGameID"`
	StatusID              uint
	Status                GameStatus `gorm:"foreignKey:StatusID"`
	Storyline             string
	Summary               string
	AggregatedRating      float64
	AggregatedRatingCount int
	Rating                float64
	RatingCount           int
	Follows               int
	Hypes                 int
	TotalRating           float64
	TotalRatingCount      int
	VersionParentID       uint
	VersionParent         *Game `gorm:"foreignKey:VersionParentID"`
	VersionTitle          string
	Cover                 Cover
	Genres                []Genre              `gorm:"many2many:game_genre;"`
	Characters            []*Character         `gorm:"many2many:game_character;"`
	AlternativeNames      []AlternativeName    `gorm:"foreignKey:GameID"`
	AgeRatings            []AgeRating          `gorm:"many2many:game_age_rating;"`
	Artworks              []Artwork            `gorm:"foreignKey:GameID"`
	Series                []*Series            `gorm:"many2many:game_series;"`
	Bundles               []*Game              `gorm:"many2many:game_bundle;"`
	DLCS                  []*Game              `gorm:"many2many:game_dlc;"`
	ExpandedGames         []*Game              `gorm:"many2many:game_expanded_game;"`
	StandaloneExpansions  []*Game              `gorm:"many2many:game_standalone_expansion;"`
	Expansions            []*Game              `gorm:"many2many:game_expansion;"`
	Forks                 []*Game              `gorm:"many2many:game_fork;"`
	Ports                 []*Game              `gorm:"many2many:game_port;"`
	Remakes               []*Game              `gorm:"many2many:game_remake;"`
	Remasters             []*Game              `gorm:"many2many:game_remaster;"`
	SimilarGames          []*Game              `gorm:"many2many:game_similar_game;"`
	ExternalGames         []ExternalGame       `gorm:"foreignKey:GameID"`
	GameFranchises        []GameFranchise      `gorm:"foreignKey:GameID"`
	GameEngines           []*GameEngine        `gorm:"many2many:game_game_engine;"`
	GameLocalizations     []GameLocalization   `gorm:"foreignKey:GameID"`
	GameModes             []GameMode           `gorm:"many2many:game_game_mode;"`
	GameCompanies         []GameCompany        `gorm:"foreignKey:GameID"`
	Keywords              []Keyword            `gorm:"many2many:game_keyword;"`
	LanguangeSupports     []LanguageSupport    `gorm:"foreignKey:GameID"`
	MultiplayerModes      []MultiplayerMode    `gorm:"foreignKey:GameID"`
	Platforms             []*Platform          `gorm:"many2many:game_platform;"`
	PlayerPerspectives    []*PlayerPerspective `gorm:"many2many:game_player_perspective;"`
	ReleaseDates          []ReleaseDate        `gorm:"foreignKey:GameID"`
	Screenshots           []Screenshot         `gorm:"foreignKey:GameID"`
	Themes                []Theme              `gorm:"many2many:game_theme;"`
	Videos                []GameVideo          `gorm:"foreignKey:GameID"`
	Websites              []Website            `gorm:"foreignKey:GameID"`
}

type GameFranchise struct {
	gorm.Model
	GameID      uint
	FranchiseID uint
	Franchise   Franchise `gorm:"foreignKey:FranchiseID"`
	isMain      bool
}
