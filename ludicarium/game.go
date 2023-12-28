package ludicarium

import (
	"github.com/vikovanesta/ludicarium-api/igdb"
)

type Game struct {
	ID                    int               `json:"id"`
	AgeRatings            []int             `json:"age_ratings"`
	AggregatedRating      float64           `json:"aggregated_rating"`
	AggregatedRatingCount int               `json:"aggregated_rating_count"`
	AlternativeNames      []int             `json:"alternative_names"`
	Artworks              []int             `json:"artworks"`
	Bundles               []int             `json:"bundles"`
	Category              igdb.GameCategory `json:"category"`
	Collection            int               `json:"collection"`
	Collections           []int             `json:"collections"`
	Cover                 int               `json:"cover"`
	CreatedAt             int               `json:"created_at"`
	DLCS                  []int             `json:"dlcs"`
	Expansions            []int             `json:"expansions"`
	ExternalGames         []int             `json:"external_games"`
	FirstReleaseDate      int               `json:"first_release_date"`
	Follows               int               `json:"follows"`
	Forks                 []int             `json:"forks"`
	Franchise             int               `json:"franchise"`
	Franchises            []int             `json:"franchises"`
	GameEngines           []int             `json:"game_engines"`
	GameLocalizations     []int             `json:"game_localizations"`
	GameModes             []int             `json:"game_modes"`
	Genres                []int             `json:"genres"`
	Hypes                 int               `json:"hypes"`
	InvolvedCompanies     []int             `json:"involved_companies"`
	Keywords              []int             `json:"keywords"`
	LanguangeSupports     []int             `json:"language_supports"`
	MultiplayerModes      []int             `json:"multiplayer_modes"`
	Name                  string            `json:"name"`
	ParentGame            int               `json:"parent_game"`
	Platforms             []int             `json:"platforms"`
	PlayerPerspectives    []int             `json:"player_perspectives"`
	Ports                 []int             `json:"ports"`
	Rating                float64           `json:"rating"`
	RatingCount           int               `json:"rating_count"`
	ReleaseDates          []int             `json:"release_dates"`
	Remakes               []int             `json:"remakes"`
	Remasters             []int             `json:"remasters"`
	Screenshots           []int             `json:"screenshots"`
	SimilarGames          []int             `json:"similar_games"`
	Slug                  string            `json:"slug"`
	StandaloneExpansions  []int             `json:"standalone_expansions"`
	Status                igdb.GameStatus   `json:"status"`
	Storyline             string            `json:"storyline"`
	Summary               string            `json:"summary"`
	Tags                  []igdb.Tag        `json:"tags"`
	Themes                []int             `json:"themes"`
	TotalRating           float64           `json:"total_rating"`
	TotalRatingCount      int               `json:"total_rating_count"`
	UpdatedAt             int               `json:"updated_at"`
	URL                   string            `json:"url"`
	VersionParent         int               `json:"version_parent"`
	VersionTitle          string            `json:"version_title"`
	Videos                []int             `json:"videos"`
	Websites              []int             `json:"websites"`
}

type GameService interface {
	Get(id int) (*Game, error)
	List(ids []int) ([]*Game, error)
	Index() ([]*Game, error)
}
