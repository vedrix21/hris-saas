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
        role, _ := c.Cookie("role")
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
		c.Set("user", models.User{
            Username: username,
            Role:     role,
        })

		c.Next()
	}
}