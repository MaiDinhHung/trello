package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	ListID      uint      `json:"list_id"`
	List        List      `gorm:"constraint:OnDelete:CASCADE;"`
	UserID      uint      `json:"user_id"`
	Comments    []Comment `json:"comments" gorm:"foreignKey:card_id"`
	Labels      []Label   `json:"labels" gorm:"many2many:card_labels;"`
	Position    int       `json:"position"`
	ID          uint      `json:"id"`
}
