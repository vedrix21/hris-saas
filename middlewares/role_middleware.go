package middlewares

import (
    "net/http"

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