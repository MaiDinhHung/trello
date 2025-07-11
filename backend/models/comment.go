package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	CardID    uint      `json:"card_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
