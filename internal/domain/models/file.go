package models

import (
	"github.com/ahmetilboga2004/internal/application/dto"
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

func (f *File) ToBasicInfoDTO() dto.FileBasicInfo {
	return dto.FileBasicInfo{
		ID:       f.ID,
		Name:     f.Name,
		FileType: f.FileType,
		Size:     f.Size,
		Public:   f.Public,
	}
}

func (f *File) ToDetailsDTO() dto.FileDetails {
	path := ""
	if f.Public {
		path = f.Path
	}

	return dto.FileDetails{
		ID:        f.ID,
		Name:      f.Name,
		Path:      path,
		FileType:  f.FileType,
		Size:      f.Size,
		Public:    f.Public,
		CreatedAt: f.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: f.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
