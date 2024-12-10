package conrollers

import (
	"bar_site/configs"
	"bar_site/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDrink создаёт новый напиток
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
