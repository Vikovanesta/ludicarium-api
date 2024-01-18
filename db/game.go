package db

import "time"

type Game struct {
	Model
	Name                  string               `gorm:"not null" json:"name"`
	Slug                  string               `gorm:"not null" json:"slug"`
	CategoryID            uint                 `json:"-"`
	Category              GameCategory         `gorm:"foreignKey:CategoryID" json:"category"`
	FirstReleaseDate      time.Time            `json:"first_release_date"`
	ParentGameID          uint                 `json:"-"`
	ParentGame            *Game                `gorm:"foreignKey:ParentGameID" json:"parent_game"`
	StatusID              uint                 `json:"-"`
	Status                GameStatus           `gorm:"foreignKey:StatusID" json:"status"`
	Storyline             string               `json:"storyline"`
	Summary               string               `json:"summary"`
	AggregatedRating      float64              `json:"aggregated_rating"`
	AggregatedRatingCount int                  `json:"aggregated_rating_count"`
	Rating                float64              `json:"rating"`
	RatingCount           int                  `json:"rating_count"`
	TotalRating           float64              `json:"total_rating"`
	TotalRatingCount      int                  `json:"total_rating_count"`
	Follows               int                  `json:"follows"`
	Hypes                 int                  `json:"hypes"`
	VersionParentID       uint                 `json:"-"`
	VersionTitle          string               `json:"version_title"`
	VersionParent         *Game                `gorm:"foreignKey:VersionParentID" json:"version_parent"`
	CoverID               uint                 `json:"-"`
	Cover                 Cover                `gorm:"foreignKey:CoverID" json:"cover"`
	Genres                []Genre              `gorm:"many2many:game_genre;" json:"genres"`
	Characters            []*Character         `gorm:"many2many:game_character;" json:"characters"`
	AlternativeNames      []AlternativeName    `gorm:"foreignKey:GameID" json:"alternative_names"`
	AgeRatings            []AgeRating          `gorm:"many2many:game_age_rating;" json:"age_ratings"`
	Artworks              []Artwork            `gorm:"foreignKey:GameID" json:"artworks"`
	Series                []*Series            `gorm:"many2many:game_series;" json:"series"`
	Bundles               []*Game              `gorm:"many2many:game_bundle;" json:"bundles"`
	DLCS                  []*Game              `gorm:"many2many:game_dlc;" json:"dlcs"`
	ExpandedGames         []*Game              `gorm:"many2many:game_expanded_game;" json:"expanded_games"`
	StandaloneExpansions  []*Game              `gorm:"many2many:game_standalone_expansion;" json:"standalone_expansions"`
	Expansions            []*Game              `gorm:"many2many:game_expansion;" json:"expansions"`
	Forks                 []*Game              `gorm:"many2many:game_fork;" json:"forks"`
	Ports                 []*Game              `gorm:"many2many:game_port;" json:"ports"`
	Remakes               []*Game              `gorm:"many2many:game_remake;" json:"remakes"`
	Remasters             []*Game              `gorm:"many2many:game_remaster;" json:"remasters"`
	SimilarGames          []*Game              `gorm:"many2many:game_similar_game;" json:"similar_games"`
	ExternalGames         []ExternalGame       `gorm:"foreignKey:GameID" json:"external_games"`
	GameFranchises        []GameFranchise      `gorm:"foreignKey:GameID" json:"game_franchises"`
	GameEngines           []*GameEngine        `gorm:"many2many:game_game_engine;" json:"game_engines"`
	GameLocalizations     []GameLocalization   `gorm:"foreignKey:GameID" json:"game_localizations"`
	GameModes             []GameMode           `gorm:"many2many:game_game_mode;" json:"game_modes"`
	RelatedCompanies      []GameCompany        `gorm:"foreignKey:GameID" json:"related_companies"`
	Keywords              []Keyword            `gorm:"many2many:game_keyword;" json:"keywords"`
	LanguangeSupports     []LanguageSupport    `gorm:"foreignKey:GameID" json:"language_supports"`
	MultiplayerModes      []MultiplayerMode    `gorm:"foreignKey:GameID" json:"multiplayer_modes"`
	Platforms             []*Platform          `gorm:"many2many:game_platform;" json:"platforms"`
	PlayerPerspectives    []*PlayerPerspective `gorm:"many2many:game_player_perspective;" json:"player_perspectives"`
	ReleaseDates          []ReleaseDate        `gorm:"foreignKey:GameID" json:"release_dates"`
	Screenshots           []Screenshot         `gorm:"foreignKey:GameID" json:"screenshots"`
	Themes                []Theme              `gorm:"many2many:game_theme;" json:"themes"`
	Videos                []GameVideo          `gorm:"foreignKey:GameID" json:"videos"`
	Websites              []Website            `gorm:"foreignKey:GameID" json:"websites"`
}
