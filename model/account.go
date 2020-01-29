package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	Enabled   bool
	Token     string
	ExpiredAt *time.Time
	Contents  []Content
}
