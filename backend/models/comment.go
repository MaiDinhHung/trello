package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string    `json:"content"`
	CardID    uint      `json:"CardID"`
	UserID    uint      
	User    User `gorm:"foreignKey:UserID"`
}
