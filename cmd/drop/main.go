package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/vikovanesta/ludicarium-api/db"
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	database.Migrator().DropTable(
		&db.DateCategory{},
		&db.WebsiteCategory{},
		&db.Region{},

		"age_rating_content_description",
		"game_age_rating",
		&db.AgeRatingCategory{},
		&db.AgeRatingContentCategory{},
		&db.AgeRatingContentDescription{},
		&db.AgeRatingTitle{},
		&db.AgeRating{},

		&db.GameCategory{},
		&db.GameMode{},
		&db.GameStatus{},
		&db.Theme{},
		&db.Genre{},
		&db.Keyword{},
		&db.PlayerPerspective{},

		"game_series",
		&db.SeriesType{},
		&db.Series{},
		&db.SeriesRelation{},

		&db.CharacterGender{},
		&db.CharacterMugShot{},
		&db.CharacterSpecies{},

		"platform_platform_version_company",
		"game_engine_company",
		&db.GameCompany{},
		&db.CompanyLogo{},
		&db.CompanyWebsite{},
		&db.Company{},

		"platform_platform_version",
		"game_engine_platform",
		"game_platform",
		&db.PlatformCategory{},
		&db.PlatformFamily{},
		&db.PlatformLogo{},
		&db.PlatformWebsite{},
		&db.PlatformVersionReleaseDate{},
		&db.PlatformVersionCompany{},
		&db.PlatformVersion{},
		&db.Platform{},

		&db.ExternalGameCategory{},
		&db.ExternalGameMedia{},

		&db.AlternativeName{},
		&db.Artwork{},
		&db.Cover{},
		&db.MultiplayerMode{},
		&db.Screenshot{},
		&db.GameVideo{},
		&db.Website{},

		&db.Franchise{},
		&db.GameFranchise{},

		"game_game_engine",
		&db.GameEngineLogo{},
		&db.GameEngine{},
		&db.GameLocalization{},

		&db.Language{},
		&db.LanguageSupportType{},
		&db.LanguageSupport{},

		"game_bundle",
		"game_character",
		"game_dlc",
		"game_expanded_game",
		"game_expansion",
		"game_fork",
		"game_game_mode",
		"game_genre",
		"game_keyword",
		"game_player_perspective",
		"game_port",
		"game_remake",
		"game_remaster",
		"game_similar_game",
		"game_standalone_expansion",
		"game_theme",
		"game_version_feature",
		"game_version_game",
		&db.ExternalGame{},
		&db.Game{},
		&db.Character{},

		&db.GameVersionFeatureCategory{},
		&db.GameVersionFeatureInclusion{},
		&db.GameVersionFeature{},
		&db.GameVersionFeatureValue{},
		&db.GameVersion{},

		"event_game",
		"event_video",
		&db.EventLogo{},
		&db.EventNetworkType{},
		&db.EventNetwork{},
		&db.Event{},

		&db.ReleaseDateStatus{},
		&db.ReleaseDate{},
	)

}
