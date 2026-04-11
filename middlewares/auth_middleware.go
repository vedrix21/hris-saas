package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        _, err := c.Cookie("user")

        if err != nil {
            c.Redirect(302, "/login")
            c.Abort()
            return
        }

        c.Next()
    }
}