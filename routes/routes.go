package routes

import (
	"hris/controllers"
	"hris/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// ===== PUBLIC =====
	r.GET("/", controllers.Home)
	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)

	r.GET("/logout", controllers.Logout)

	r.GET("/forgot-password", controllers.ShowForgotPassword)
	r.POST("/forgot-password", controllers.ForgotPassword)
	r.GET("/reset-password", controllers.ShowResetPassword)
	r.POST("/reset-password", controllers.ResetPassword)

	// ===== AUTH (SEMUA USER LOGIN) =====
	auth := r.Group("/")

	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/dashboard", controllers.Dashboard)
		auth.GET("/switch-env/:env", controllers.SwitchEnv)
	}

	// ===== OWNER ONLY =====
	owner := r.Group("/owner")
	owner.Use(middlewares.AuthMiddleware())
	owner.Use(middlewares.OwnerOnly())
	{
		owner.GET("/dashboard", controllers.OwnerDashboard)
		owner.GET("/create-account", controllers.ShowCreateAccount)
		owner.GET("/settings", controllers.ShowSettings)
		owner.POST("/settings/create-account", controllers.CreateAccount)
	}
}
