package db

type Franchise struct {
	Model
	Name           string          `gorm:"not null" json:"name"`
	Slug           string          `gorm:"not null" json:"slug"`
	GameFranchises []GameFranchise `gorm:"foreignKey:FranchiseID" json:"gameFranchises"`
}
