package db

type AgeRatingContentDescription struct {
	Model
	CategoryID  uint                     `json:"-"`
	Category    AgeRatingContentCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Description string                   `json:"description"`
}
