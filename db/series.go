package db

type Series struct {
	Model
	Name         string           `gorm:"not null" json:"name"`
	Slug         string           `gorm:"not null" json:"slug"`
	TypeID       uint             `json:"-"`
	Type         SeriesType       `gorm:"foreignKey:TypeID" json:"type"`
	ChildSeries  []SeriesRelation `gorm:"foreignKey:ParentSeriesID" json:"child_series"`
	ParentSeries []SeriesRelation `gorm:"foreignKey:ChildSeriesID" json:"parent_series"`
	Game         []*Game          `gorm:"many2many:game_series;" json:"games"`
}
