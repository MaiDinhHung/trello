package models

import (
	"gorm.io/gorm"
)


type ChecklistItem struct {
	gorm.Model
	ID      uint   `json:"id"`
	Content   string         `json:"content"`
	Completed bool           `json:"completed"`
	CardID    uint           `json:"card_id"`
}