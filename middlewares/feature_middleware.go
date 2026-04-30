package middlewares

import (
	"hris/models"
	"hris/services"

	"github.com/gin-gonic/gin"
)

func RequireFeature(feature string) gin.HandlerFunc {
	return func(c *gin.Context) {

		account := c.MustGet("account").(models.Account)

		if !services.HasFeature(account, feature) {
			c.String(403, "Upgrade plan to access this feature 🚀")
			c.Abort()
			return
		}

		c.Next()
	}
}