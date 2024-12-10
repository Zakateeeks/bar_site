package initial

import (
	"bar_site/configs"
	"bar_site/internal/models"
)

// SyncDatabase Функция для автомаической миграции данных
func SyncDatabase() {
	err := configs.DB.AutoMigrate(&models.DrinkModel{})
	if err != nil {
		panic("failed to migrate db")
		return
	}
}
