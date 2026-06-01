package middlewares

import (
    "net/http"
    "hris/models"

    "github.com/gin-gonic/gin"
)

// 🔐 Middleware untuk cek role owner
func OwnerOnly() gin.HandlerFunc {
    return func(c *gin.Context) {

        role, err := c.Cookie("role")
        if err != nil || role != "owner" {
            c.String(http.StatusForbidden, "Forbidden - Owner only")
            c.Abort()
            return
        }

        c.Next()
    }
}

func SuperAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		user := c.MustGet("user").(models.User)

		if user.Role != "superadmin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})
			return
		}

		c.Next()
	}
}