package conrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDrink создаёт новый напиток
func CreateDrink(c *gin.Context) {
	var drink struct {
		Title string
		Type  string
		Brand string
		Count int
		Image string
	}

	if c.Bind(&drink) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read drink"})
		return
	}
}

func GetDrink() gin.HandlerFunc {
	//ToDO
}
