package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {

        user, err := c.Cookie("user")
        if err != nil || user == "" {
            c.Redirect(http.StatusFound, "/login")
            c.Abort()
            return
        }

        c.Next()
    }
}