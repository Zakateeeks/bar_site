package conrollers

import (
	"bar_site/configs"
	"bar_site/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDrink(c *gin.Context) {
	var drink models.DrinkModel

	if c.Bind(&drink) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read drink"})
		return
	}

	if err := configs.DB.Create(&drink).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create drink1"})
		return
	}

}

func GetDrink() gin.HandlerFunc {
	return func(c *gin.Context) {
		var drinks []models.DrinkModel
		if err := configs.DB.Find(&drinks).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get drinks"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"drinks": drinks})
	}
}

func UpdateDrink(c *gin.Context) {
	var requestBody struct {
		DrinkId int `json:"drink_id"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	var drink models.DrinkModel
	if err := configs.DB.Where("id = ?", requestBody.DrinkId).First(&drink).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drink not found"})
		return
	}
	if drink.Count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Drink out of stock"})
		return
	}
	drink.Count -= 1
	if err := configs.DB.Save(&drink).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update drink"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"drink": drink})
}
