package model

import "github.com/jinzhu/gorm"

type Place struct {
	gorm.Model
	Lat      string    `validate:"required,lat"`
	Lon      string    `validate:"required,lon"`
	Contents []Content `gorm:"foreignkey:PlaceID"`
}
