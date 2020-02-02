package model

import (
	"github.com/jinzhu/gorm"
)

type ContentStatus int

const (
	Draft     ContentStatus = 0
	Published ContentStatus = 1
	Archived  ContentStatus = 2
)

type Content struct {
	gorm.Model
	Title       string `gorm:"size:512; index:contents_title"`
	Slug        string `gorm:"type:varchar(512); unique_index"`
	Description string
	PlaceID     uint          `gorm:"index:contents_place_id"`
	UserID      uint          `gorm:"index:contents_author_id"`
	Status      ContentStatus `sql:"DEFAULT: 0"`
}

// Archive, sets archive status as `Published`
func (c *Content) Publish() {
	c.Status = Published
}

// Archive, sets archive status as `Archived`
func (c *Content) Archive() {
	c.Status = Archived
}
