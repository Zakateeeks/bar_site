package routes

import (
	"bar_site/internal/conrollers"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(router *gin.Engine) {
	router.GET("/", conrollers.HomePage)
	router.GET("/drinks", conrollers.LookDrinks)
}
