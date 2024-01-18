package ludicarium

import (
	"errors"
	"time"

	"github.com/vikovanesta/ludicarium-api/db"
	"github.com/vikovanesta/ludicarium-api/igdb"
)

type GameService struct {
	db   *db.DB
	igdb *igdb.Client
}

func NewGameService(db *db.DB, igdb *igdb.Client) *GameService {
	return &GameService{
		db:   db,
		igdb: igdb,
	}
}

func (gs *GameService) List() ([]*igdb.Game, error) {
	var igdbGames []*igdb.Game
	var err error

	igdbGames, err = gs.igdb.Games.Index(igdb.SetLimit(15), igdb.SetFields("*"))
	if err != nil {
		return nil, err
	}

	return igdbGames, nil
}

func (gs *GameService) Create(game interface{}) (*db.Game, error) {
	switch g := game.(type) {
	case *db.Game:
		result := gs.db.Where(db.Game{Name: g.Name}).Assign(g).FirstOrCreate(g)
		if result.Error != nil {
			return nil, result.Error
		}
		return g, nil
	case *igdb.Game:
		dbGame := igdbGameToLudicariumGame(g)
		result := gs.db.Where(db.Game{Name: dbGame.Name}).Assign(dbGame).FirstOrCreate(dbGame)
		if result.Error != nil {
			return nil, result.Error
		}
		return dbGame, nil
	default:
		return nil, errors.New("unsupported game type")
	}
}

func igdbGameToLudicariumGame(g *igdb.Game) *db.Game {
	return &db.Game{
		Model: db.Model{
			ID:        uint(g.ID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:                  g.Name,
		Slug:                  g.Slug,
		CategoryID:            uint(g.Category),
		FirstReleaseDate:      time.Unix(int64(g.FirstReleaseDate), 0),
		ParentGameID:          uint(g.ParentGame),
		StatusID:              uint(g.Status),
		Storyline:             g.Storyline,
		Summary:               g.Summary,
		AggregatedRating:      g.AggregatedRating,
		AggregatedRatingCount: g.AggregatedRatingCount,
		Rating:                g.Rating,
		RatingCount:           g.RatingCount,
		TotalRating:           g.TotalRating,
		TotalRatingCount:      g.TotalRatingCount,
		Follows:               g.Follows,
		Hypes:                 g.Hypes,
		VersionParentID:       uint(g.VersionParent),
		VersionTitle:          g.VersionTitle,
		CoverID:               uint(g.Cover),
	}
}
