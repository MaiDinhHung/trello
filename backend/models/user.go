package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Password string    `json:"password"`
	Boards   []Board   `gorm:"foreignKey:user_id"`
	Cards    []Card    `gorm:"foreignKey:user_id"`
	Comments []Comment `gorm:"foreignKey:user_id"`
}
