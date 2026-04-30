package routes

import (
	"hris/controllers"
	"hris/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// =============================
	// 🔓 PUBLIC (NO LOGIN)
	// =============================
	r.GET("/", controllers.Home)

	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)

	// 🔥 BEST PRACTICE: logout pakai POST
	r.POST("/logout", controllers.Logout)

	r.GET("/forgot-password", controllers.ShowForgotPassword)
	r.POST("/forgot-password", controllers.ForgotPassword)

	r.GET("/reset-password", controllers.ShowResetPassword)
	r.POST("/reset-password", controllers.ResetPassword)

	// =============================
	// 🔐 AUTH (LOGIN REQUIRED)
	// =============================
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		// Dashboard
		auth.GET("/dashboard", controllers.Dashboard)

		// Env switch
		// auth.GET("/switch-env/:env", controllers.SwitchEnv)

		auth.GET("/employees",
			middlewares.RequireFeature("employee"),
			controllers.Employees,
		)

		auth.GET("/attendance",
			middlewares.RequireFeature("attendance"),
			controllers.AttendancePage,
		)

		auth.GET("/payroll",
			middlewares.RequireFeature("payroll"),
			controllers.PayrollPage,
		)

		// =============================
		// 💳 BILLING / SUBSCRIPTION
		// =============================
		auth.POST("/billing/upgrade", controllers.UpgradePlan)
	}

	// =============================
	// 👑 OWNER (SUPER ADMIN)
	// =============================
	owner := r.Group("/owner")
	owner.Use(middlewares.AuthMiddleware())
	owner.Use(middlewares.OwnerOnly())
	{
		owner.GET("/dashboard", controllers.OwnerDashboard)

		owner.GET("/create_account", controllers.ShowCreateAccount)
		owner.POST("/create_account", controllers.CreateAccount)

		owner.GET("/settings", controllers.ShowSettings)
	}
}
