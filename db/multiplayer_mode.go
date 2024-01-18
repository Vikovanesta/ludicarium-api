package db

type MultiplayerMode struct {
	Model
	Campaigncoop      bool     `json:"campaign_coop"`
	Dropin            bool     `json:"drop_in"`
	GameID            uint     `gorm:"not null" json:"-"`
	Lancoop           bool     `json:"lan_coop"`
	Offlinecoop       bool     `json:"offline_coop"`
	Offlinecoopmax    int      `json:"offline_coop_max"`
	Offlinemax        int      `json:"offline_max"`
	Onlinecoop        bool     `json:"online_coop"`
	Onlinecoopmax     int      `json:"online_coop_max"`
	Onlinemax         int      `json:"online_max"`
	PlatformID        uint     `json:"-"`
	Platform          Platform `gorm:"foreignKey:PlatformID" json:"platform"`
	Splitscreen       bool     `json:"splitscreen"`
	Splitscreenonline bool     `json:"splitscreen_online"`
}
