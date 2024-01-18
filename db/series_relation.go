package db

type SeriesRelation struct {
	Model
	ChildSeriesID  uint   `gorm:"not null" json:"-"`
	ChildSeries    Series `gorm:"foreignKey:ChildSeriesID" json:"child_series"`
	ParentSeriesID uint   `gorm:"not null" json:"-"`
	ParentSeries   Series `gorm:"foreignKey:ParentSeriesID" json:"parent_series"`
}
