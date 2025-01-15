package routes

import (
	"bar_site/internal/conrollers"
	"bar_site/internal/middleware"
	"github.com/gin-gonic/gin"
)

func BarRoute(router *gin.Engine) {
	router.POST("/drink", conrollers.CreateDrink)
	router.GET("/api/drinks", conrollers.GetDrink())
	router.POST("/api/signup", conrollers.SignUp)
	router.POST("api/login", conrollers.SignIn)
	router.POST("/api/order", middleware.AuthMiddleware(), conrollers.UpdateDrink)
}
