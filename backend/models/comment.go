package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string    `json:"content"`
	CardID    uint      `json:"card_id"`
	UserID    uint      
	User    User `gorm:"foreignKey:UserID"`
}
