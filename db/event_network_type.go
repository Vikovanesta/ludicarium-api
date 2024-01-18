package db

type EventNetworkType struct {
	Model
	Name string `gorm:"unique;not null" json:"name"`
}
