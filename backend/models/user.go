package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	Boards   []Board   `gorm:"foreignKey:UserID"`
	Cards    []Card    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:ID"`
	Members  []CardMember `gorm:"foreignKey:UserID"`
}
