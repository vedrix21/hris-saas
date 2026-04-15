package utils

import "github.com/gin-gonic/gin"

func GetEnv(c *gin.Context) string {
    env, err := c.Cookie("env")
    if err != nil {
        return "prod"
    }
    return env
}