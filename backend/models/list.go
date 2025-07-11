package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	BoardID uint   `json:"board_id"`
	Cards   []Card `json:"cards" gorm:"foreignKey:list_id;constraint:OnDelete:CASCADE"`
}
