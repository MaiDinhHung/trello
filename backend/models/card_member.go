package models

import (
	"gorm.io/gorm"
)

type CardMember struct {
	gorm.Model
	CardID uint   `json:"card_id"`
	UserID uint   `json:"UserID"`
	Role   string `json:"role"`
	User   User   `gorm:"foreignKey:UserID"`
}	