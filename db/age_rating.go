package db

type AgeRating struct {
	Model
	CategoryID          uint                          `json:"-"`
	Category            AgeRatingCategory             `gorm:"foreignKey:CategoryID" json:"category"`
	ContentDescriptions []AgeRatingContentDescription `gorm:"many2many:age_rating_content_description;" json:"content_descriptions"`
	RatingID            uint                          `json:"-"`
	Rating              AgeRatingTitle                `gorm:"foreignKey:RatingID" json:"rating"`
	Synopsis            string                        `json:"synopsis"`
	RatingCoverURL      string                        `json:"rating_cover_url"`
}
