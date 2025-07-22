package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	StartDate      string          `json:"start_date"`
	EndDate        string          `json:"end_date"`
	ListID         uint            `json:"list_id"`
	UserID         uint            `json:"UserID"`
	Comments       []Comment       `json:"comments" gorm:"foreignKey:CardID"`
	Labels         []Label         `json:"labels" gorm:"many2many:card_labels;"`
	ChecklistItems []ChecklistItem `json:"checklist_items" gorm:"foreignKey:CardID"`
	Position       int             `json:"position"`
	Members         []CardMember    `json:"members" gorm:"foreignKey:CardID"`
	Completed      bool            `json:"completed"`
}
