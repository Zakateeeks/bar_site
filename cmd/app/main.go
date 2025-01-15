package main

import (
	"bar_site/configs"
	"bar_site/internal/initial"
	"bar_site/internal/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnvVar()
	configs.ConnectToDb()
	initial.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.Static("/images", "../../web/image")
	router.LoadHTMLGlob("../../web/templates/*")
	routes.BarRoute(router)
	routes.RegisterWebRoutes(router)
	router.Run(":8080")
}
