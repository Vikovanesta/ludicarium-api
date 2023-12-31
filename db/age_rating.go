package db

import (
	"gorm.io/gorm"
)

type AgeRating struct {
	gorm.Model
	CategoryID          uint
	Category            AgeRatingCategory             `gorm:"foreignKey:CategoryID"`
	ContentDescriptions []AgeRatingContentDescription `gorm:"many2many:age_rating_content_description;"`
	RatingID            uint
	Rating              AgeRatingTitle `gorm:"foreignKey:RatingID"`
	Synopsis            string
	RatingCoverURL      string
}
