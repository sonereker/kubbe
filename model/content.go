package model

import (
	"github.com/jinzhu/gorm"
)

type Content struct {
	gorm.Model
	Title       string
	Slug        string
	Description string
	PlaceID     uint
	AuthorID    uint
	Status      int
}
