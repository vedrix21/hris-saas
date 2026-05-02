package middlewares

import (
	"hris/config"
	"hris/models"
	"hris/utils"

	"github.com/gin-gonic/gin"
)

func SubscriptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		accountCode, _ := c.Cookie("tenant")

		var acc models.Account
		config.DB.Where("code = ?", accountCode).First(&acc)

		// 🔥 CEK SUBSCRIPTION
		if !utils.IsSubscriptionActive(acc) {

			// biar gak loop ke billing
			if c.Request.URL.Path != "/billing" {
				c.Redirect(302, "/billing")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}