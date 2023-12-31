package db

type Image struct {
	ImageID      string
	AlphaChannel bool
	Animated     bool
	Height       int
	URL          string
	Width        int
}
