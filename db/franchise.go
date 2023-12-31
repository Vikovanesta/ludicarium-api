package db

import "gorm.io/gorm"

type Franchise struct {
	gorm.Model
	Name           string
	Slug           string
	GameFranchises []GameFranchise `gorm:"foreignKey:FranchiseID"`
}
