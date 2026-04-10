
package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "hris/config"
    "hris/routes"
)

func main() {
     r := gin.Default()

    // 🔥 TAMBAHKAN INI
    r.Static("/ui", "./ui")

    config.ConnectMasterDB()
    routes.SetupRoutes(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}
