package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	StartDate      string          `json:"start_date"`
	EndDate        string          `json:"end_date"`
	ListID         uint            `json:"list_id"`
	// List           List            `gorm:"constraint:OnDelete:CASCADE;"`
	UserID         uint            `json:"UserID"`
	Comments       []Comment       `json:"comments" gorm:"foreignKey:card_id"`
	Labels         []Label         `json:"labels" gorm:"many2many:card_labels;"`
	ChecklistItems []ChecklistItem `json:"checklist_items" gorm:"foreignKey:card_id"`
	Position       int             `json:"position"`
	Member         []CardMember    `json:"members" gorm:"foreignKey:card_id"`
	Completed      bool            `json:"completed"`
}
