package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, err := c.Cookie("user")
        if err != nil || user == "" {
            c.Redirect(http.StatusFound, "/")
            c.Abort()
            return
        }

        c.Next()
    }
}