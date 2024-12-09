package main

import (
	"bar_site/configs"
	"bar_site/internal/initial"
	"bar_site/internal/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnvVar()
}

func main() {
	router := gin.Default()

	configs.ConnectToDb()

	router.LoadHTMLGlob("../../templates/*")
	routes.BarRoute(router)
	routes.RegisterWebRoutes(router)
	router.Run(":8080")
}
