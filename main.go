package main

import (
	"hris/config"
	"hris/routes"
	"hris/utils"
	"hris/services"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	// 🔥 static file (css/js kalau ada)
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")

	// 🔥 load template HTML (WAJIB untuk login & dashboard)
	// r.LoadHTMLGlob("templates/**/*")

	

	config.ConnectMasterDB()
	utils.SeedOwner()
	routes.SetupRoutes(r)

	go func() {
		time.Sleep(10 * time.Second) // biar server ready dulu

		for {
			services.CheckSubscriptions()
			time.Sleep(24 * time.Hour)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	

	r.Run(":" + port)

	
}
