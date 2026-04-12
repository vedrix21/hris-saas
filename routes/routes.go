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
    
    r.GET("/owner/dashboard", controllers.OwnerDashboard)
    r.GET("/logout", controllers.Logout)

    auth := r.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    {
        auth.GET("/dashboard", controllers.Dashboard)
    }

    owner := r.Group("/owner")
    owner.Use(middlewares.AuthMiddleware())
    owner.Use(middlewares.OwnerOnly())
    {
        owner.GET("/dashboard", controllers.OwnerDashboard)
        owner.GET("/create-account", controllers.ShowCreateAccount)
        owner.POST("/create-account", controllers.CreateAccount)
    }
}

