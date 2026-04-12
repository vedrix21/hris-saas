package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func OwnerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tenant, _ := c.Cookie("tenant")

        // 🔥 owner khusus (contoh: aitherhr)
        if tenant != "aitherhr" {
            c.String(http.StatusForbidden, "Forbidden")
            c.Abort()
            return
        }

        c.Next()
    }
}