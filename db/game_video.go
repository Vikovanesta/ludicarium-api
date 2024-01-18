package db

type GameVideo struct {
	Model
	Name    string `json:"name"`
	GameID  uint   `gorm:"not null" json:"-"`
	VideoID string `gorm:"not null" json:"video_id"` // External ID for the video on the video platform
}
