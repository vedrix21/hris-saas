package middlewares

import (
	"net/http"

	"hris/services"

	"github.com/gin-gonic/gin"
)

func RequireFeature(feature string) gin.HandlerFunc {
	return func(c *gin.Context) {

		tenant, err := c.Cookie("tenant")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Tenant not found",
			})
			return
		}

		if !services.HasFeature(tenant, feature) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Upgrade plan to access this feature",
			})
			return
		}

		c.Next()
	}
}
