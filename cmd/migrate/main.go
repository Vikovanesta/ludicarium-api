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

	database.AutoMigrate(
		&db.DateCategory{},
		&db.WebsiteCategory{},
		&db.Region{},

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

		&db.SeriesType{},
		&db.Series{},
		&db.SeriesRelation{},

		&db.CharacterGender{},
		&db.CharacterMugShot{},
		&db.CharacterSpecies{},

		&db.GameCompany{},
		&db.CompanyLogo{},
		&db.CompanyWebsite{},
		&db.Company{},

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

		&db.GameEngineLogo{},
		&db.GameEngine{},
		&db.GameLocalization{},

		&db.Language{},
		&db.LanguageSupportType{},
		&db.LanguageSupport{},

		&db.ExternalGame{},
		&db.Game{},
		&db.Character{},

		&db.GameVersionFeatureCategory{},
		&db.GameVersionFeatureInclusion{},
		&db.GameVersionFeature{},
		&db.GameVersionFeatureValue{},
		&db.GameVersion{},

		&db.EventLogo{},
		&db.EventNetworkType{},
		&db.EventNetwork{},
		&db.Event{},

		&db.ReleaseDateStatus{},
		&db.ReleaseDate{},
	)

}
