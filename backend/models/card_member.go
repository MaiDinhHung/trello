package models

import (
	"gorm.io/gorm"
)

type CardMember struct {
	gorm.Model
	CardID uint   `json:"CardID"`
	UserID uint   `json:"UserID"`
	Role   string `json:"role"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
