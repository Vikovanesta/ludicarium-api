package db

import "gorm.io/gorm"

type SeriesRelation struct {
	gorm.Model
	ChildSeriesID  uint
	ChildSeries    Series `gorm:"foreignKey:ChildSeriesID"`
	ParentSeriesID uint
	ParentSeries   Series `gorm:"foreignKey:ParentSeriesID"`
}
