package routes

import (
	"bar_site/internal/conrollers"
	"github.com/gin-gonic/gin"
)

func BarRoute(router *gin.Engine) {
	router.POST("/drink", conrollers.CreateDrink)
	router.GET("/api/drinks", conrollers.GetDrink())
}
