package routes

import (
    "hris/controllers"
    "hris/middlewares"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

    

    r.GET("/", controllers.ShowLogin)
    r.GET("/login", controllers.ShowLogin)
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    {
        auth.GET("/dashboard", controllers.Dashboard)
    }
}

