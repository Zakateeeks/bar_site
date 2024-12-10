package models

import (
	"gorm.io/gorm"
)

type DrinkModel struct {
	gorm.Model
	TitleD string `json:"title_d" gorm:"not null"`
	Type   string `json:"type" gorm:"foreignKey:TypeID"`
	Brand  string `json:"brand" gorm:"foreignKey:BrandID"`
	Count  int    `json:"count" gorm:"not null"`
	Image  string `json:"image" gorm:"not null"`
}

func (DrinkModel) TableName() string {
	return "drinks"
}
