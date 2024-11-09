package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary key"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Path       string    `gorm:"type:text;not null"`
	Size       int64     `gorm:"not null"`
	FileType   string    `gorm:"type:varchar(50);not null"`
	Public     bool      `gorm:"default:false"`
	UploadedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	User       User      `gorm:"foreignKey:UserID"`
}
