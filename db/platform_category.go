package db

type PlatformCategory struct {
	Model
	Name string `gorm:";not null;unique" json:"name"`
}
