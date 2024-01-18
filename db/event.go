package db

type Event struct {
	Model
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	Description   string         `json:"description"`
	StartTime     int            `json:"start_time"`
	EndTime       int            `json:"end_time"`
	TimeZone      string         `json:"time_zone"`
	EventNetworks []EventNetwork `gorm:"foreignKey:EventID" json:"event_networks"`
	Games         []Game         `gorm:"many2many:event_game;" json:"games"`
	LogoID        uint           `json:"-"`
	Logo          EventLogo      `gorm:"foreignKey:LogoID" json:"logo"`
	LiveStreamURL string         `json:"live_stream_url"`
	Videos        []GameVideo    `gorm:"many2many:event_video;" json:"videos"`
}
