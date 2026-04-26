package middlewares

import (
	"net/http"

	"hris/config"
	"hris/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("user")
		if err != nil || username == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// 🔥 Ambil user dari database
		var user models.User
		if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// 🔥 SET ke context (INI YANG PENTING)
		c.Set("user", user)

		c.Next()
	}
}