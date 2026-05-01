package main

import (
	"hris/config"
	"hris/routes"
	"hris/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	// 🔥 static file (css/js kalau ada)
	r.Static("/static", "./static")

	// 🔥 load template HTML (WAJIB untuk login & dashboard)
	// r.LoadHTMLGlob("templates/**/*")

	config.ConnectMasterDB()
	utils.SeedOwner()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
