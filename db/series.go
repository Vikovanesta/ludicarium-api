package db

import "gorm.io/gorm"

type Series struct {
	gorm.Model
	Name         string
	Slug         string
	TypeID       uint
	Type         SeriesType       `gorm:"foreignKey:TypeID"`
	ChildSeries  []SeriesRelation `gorm:"foreignKey:ParentSeriesID"`
	ParentSeries []SeriesRelation `gorm:"foreignKey:ChildSeriesID"`
	Game         []*Game          `gorm:"many2many:game_series;"`
}
