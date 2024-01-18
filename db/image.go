package db

type Image struct {
	ImageID      string `json:"-"`
	AlphaChannel bool   `json:"alpha_channel"`
	Animated     bool   `json:"animated"`
	Height       int    `json:"height"`
	URL          string `json:"url"`
	Width        int    `json:"width"`
}
