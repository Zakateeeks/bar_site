package conrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcom to the bar site",
	})
}

func LookDrinks(c *gin.Context) {
	c.HTML(http.StatusOK, "drinks.html", gin.H{
		"title": "Choose your drink",
	})
}

func Registration(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title": "Sign up",
	})
}
