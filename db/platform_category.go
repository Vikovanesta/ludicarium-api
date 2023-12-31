package db

import "gorm.io/gorm"

type PlatformCategory struct {
	gorm.Model
	Name string `gorm:";not null;unique"`
}
