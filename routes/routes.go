package routes

import (
	"hris/controllers"
	"hris/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// =============================
	// 🔓 PUBLIC
	// =============================
	r.GET("/", controllers.Home)

	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.Login)

	r.POST("/logout", controllers.Logout)

	r.GET("/forgot-password", controllers.ShowForgotPassword)
	r.POST("/forgot-password", controllers.ForgotPassword)

	r.GET("/reset-password", controllers.ShowResetPassword)
	r.POST("/reset-password", controllers.ResetPassword)

	// =============================
	// 🔐 AUTH
	// =============================
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.Use(middlewares.SubscriptionMiddleware())
	{
		auth.GET("/dashboard", controllers.Dashboard)

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

		// 🔥 billing WAJIB ADA
		auth.GET("/billing", controllers.BillingPage)
		auth.POST("/billing/upgrade", controllers.UpgradePlan)
		auth.POST("/billing/upload", controllers.UploadPayment)
	}

	// =============================
	// 👑 OWNER
	// =============================
	owner := r.Group("/owner")
	owner.Use(middlewares.AuthMiddleware())
	owner.Use(middlewares.OwnerOnly())
	{
		owner.GET("/dashboard", controllers.OwnerDashboard)

		owner.GET("/create_account", controllers.ShowCreateAccount)
		owner.POST("/create_account", controllers.CreateAccount)

		owner.GET("/settings", controllers.ShowSettings)

		owner.GET("/payments", controllers.PaymentList)
		owner.POST("/approve-payment", controllers.ApprovePayment)
	}
}
