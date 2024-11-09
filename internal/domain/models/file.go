package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Path     string `gorm:"type:text;not null"`
	Size     int64  `gorm:"not null"`
	FileType string `gorm:"type:varchar(50);not null"`
	Public   bool   `gorm:"default:false"`
	UserID   uint
	User     User
}
