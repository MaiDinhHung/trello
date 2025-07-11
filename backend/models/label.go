	package models

	import "gorm.io/gorm"

	type Label struct {
		gorm.Model
		Name  string `json:"name"`
		Color string `json:"color"`
		Card  []Card `json:"cards" gorm:"many2many:card_labels;"`
	}
