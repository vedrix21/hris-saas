package controllers

import (
    "github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
    tenant, _ := c.Cookie("tenant")

    c.HTML(200, "dashboard.html", gin.H{
        "tenant": tenant,
        "totalEmployee": 120, // 🔥 nanti diganti dari DB
    })
}