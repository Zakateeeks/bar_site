package models

import (
	"gorm.io/gorm"
)

type DrinkModel struct {
	gorm.Model
	Title string     `json:"title" gorm:"not null"`
	Type  TypeDrink  `json:"type" gorm:"foreignKey:TypeID"`
	Brand BrandDrink `json:"brand" gorm:"foreignKey:BrandID"`
	Count int        `json:"count" gorm:"not null"`
	Image string     `json:"image" gorm:"not null"`
}
