package db

type EventNetwork struct {
	Model
	EventID       uint             `gorm:"not null" json:"-"`
	Event         Event            `gorm:"foreignKey:EventID" json:"event"`
	NetworkTypeID uint             `gorm:"not null" json:"-"`
	NetworkType   EventNetworkType `gorm:"foreignKey:NetworkTypeID" json:"network_type"`
	URL           string           `json:"url"`
}
