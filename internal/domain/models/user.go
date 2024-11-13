package models

import (
	"github.com/ahmetilboga2004/internal/application/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Username  string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Files     []File
}

func (u *User) ToBasicInfoDTO() dto.UserBasicInfo {
	return dto.UserBasicInfo{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
	}
}

func (u *User) ToDetailsDTO() dto.UserDetails {
	return dto.UserDetails{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
	}
}
