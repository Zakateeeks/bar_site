package models

import "gorm.io/gorm"

// User Структура пользователя в БД
type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"pass" gorm:"not null"`
	RefreshToken string `json:"refresh_token"`
}
