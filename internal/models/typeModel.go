package models

import "gorm.io/gorm"

type TypeDrink struct {
	gorm.Model
	Name string `json:"tName" gorm:"not null"`
}
