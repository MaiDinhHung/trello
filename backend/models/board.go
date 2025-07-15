package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Title  string `json:"title"`
	UserID uint   `json:"UserID"`
	Lists  []List `json:"lists" gorm:"foreignKey:board_id"`
}
