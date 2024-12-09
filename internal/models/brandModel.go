package models

import "gorm.io/gorm"

type BrandDrink struct {
	gorm.Model
	Name string `json:"bName" gorm:"not null"`
}
