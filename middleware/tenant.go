
package middleware

import (
    "strings"
    "github.com/gin-gonic/gin"
    "hris/config"
)

func TenantMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        host := c.Request.Host
        sub := strings.Split(host, ".")[0]

        dsn := "saas_user:password@tcp(127.0.0.1:3306)/hris_" + sub + "?charset=utf8mb4&parseTime=True&loc=Local"

        db, _ := config.ConnectTenantDB(dsn)

        c.Set("tenantDB", db)
        c.Next()
    }
}
