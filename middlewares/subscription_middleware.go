package middlewares

import (
	"hris/config"
	"hris/models"
	"hris/utils"
	"strings"
	

	"github.com/gin-gonic/gin"
)

func SubscriptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		accountCode, _ := c.Cookie("tenant")

		var acc models.Account
		config.DB.Where("code = ?", accountCode).First(&acc)

		path := c.Request.URL.Path

		if !utils.IsSubscriptionActive(acc) {

			// 🔥 allow semua billing route
			if strings.HasPrefix(path, "/billing") {
				c.Next()
				return
			}

			c.Redirect(302, "/billing")
			c.Abort()
			return
		}

		c.Next()
	}
}